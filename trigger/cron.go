package trigger

import (
	"errors"
	"time"
)

//一分钟代表一个key，一小时60分钟 也就是说同一时间最多60个key
type Trigger struct {
	lists        [61]*TaskList //任务队列
	runListIndex int           //当前运行列表索引
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

func (c Trigger) AddTask(t Task) error {
	if t.Job == nil {
		return ErrorsInvalidTask
	}
	e := &item{task: t}
	list := c.lists[t.time.Minute()]
	//list := c.lists[1]
	if t.time.Minute() == c.runListIndex && c.runListIndex != 0 {
		list.ch <- e
	} else {
		list.insert(e)
	}
	return nil
}

func (c Trigger) Start() {
	for {
		now := time.Now()
		//listIndex := now.Minute()
		//c.runListIndex = listIndex
		c.runListIndex = now.Minute()
		list := c.lists[c.runListIndex]
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
		}

	}
}
