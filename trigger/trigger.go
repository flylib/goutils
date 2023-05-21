package trigger

import (
	"errors"
	"strconv"
	"time"
)

const (
	timeFormatKey = "2006010215"
)

//一分钟代表一个key，一小时60分钟 也就是说同一时间最多60个key
type Trigger struct {
	lists        [61]*TaskList      //当前执行任务队列
	runListIndex int                //当前运行列表索引
	waitLists    map[string][]*item //其他等待运行队列 year_moth_day_hour:[]*item
}

type TaskList struct {
	ch chan *item
	List
}

type Job interface {
	Run()
}

type Task struct {
	time time.Time //事件触发时间
	Job
}

var (
	ErrorsInvalidTask = errors.New("invalid task")
)

func NewTrigger() Trigger {
	new := Trigger{}
	for i := 0; i < len(new.lists); i++ {
		new.lists[i] = &TaskList{
			ch:   make(chan *item),
			List: newList(),
		}
	}
	return new
}

func (t Trigger) AddTask(job Job, triggerAt time.Time) error {
	if job == nil {
		return ErrorsInvalidTask
	}
	newTask := Task{
		triggerAt,
		job,
	}
	newItem := &item{task: newTask}
	now := time.Now()
	key := now.Format(timeFormatKey)
	triggerAtKey := triggerAt.Format(timeFormatKey)
	nowHourUnix, _ := strconv.Atoi(key)
	triggerAtHourUnix, _ := strconv.Atoi(triggerAtKey)
	switch {
	case nowHourUnix > triggerAtHourUnix:
		go job.Run() //立即触发
	case nowHourUnix == triggerAtHourUnix:
		minute := now.Minute()
		list := t.lists[minute]
		if newTask.time.Minute() == t.runListIndex {
			list.ch <- newItem
		} else {
			list.insert(newItem)
		}
	default:
		t.waitLists[triggerAtKey] = append(t.waitLists[triggerAtKey], newItem)
	}
	return nil
}

func (t Trigger) Start() {
	go func() {
		for {
			now := time.Now()
			t.runListIndex = 1
			preList := t.lists[t.runListIndex]
			for preList.len > 0 {
				item := preList.pop()
				go item.task.Job.Run()
			}
			list := t.lists[t.runListIndex]
			item := list.Front()
			var timer *time.Timer
			if item != nil {
				sub := item.task.time.Sub(now)
				timer = time.NewTimer(sub)
			} else {
				timer = time.NewTimer(time.Duration(60 - now.Second()))
			}
			for {
				select {
				case <-timer.C:
					if item != nil {
						go item.task.Run()
					}
					list.pop()
					break
				case newItem := <-list.ch:
					if newItem.task.time.Before(item.task.time) {
						timer.Stop()
						break
					}
					list.insert(newItem)
				}
				break
			}
		}
	}()

}
