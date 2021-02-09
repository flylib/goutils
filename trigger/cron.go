package trigger

import "time"

//一分钟代表一个key，一小时60分钟 也就是说同一时间最多60个key
type Trigger struct {
	mListJob map[int64]Job
	workers  [60]worker //工作者
}

type Job interface {
	Run()
}

type task struct {
	At time.Time //事件触发时间
	Job
}

const (
	worker_status_stop int8 = 0
	worker_status_run  int8 = 1
)

type worker struct {
	status int8 //状态
	events []task
}

func (c Trigger) Start() {
	for i := 0; i < 60; i++ {
		go func() {
			for {
				now := time.Now()
				//如果没有任务就退出当前routine
				if len(c.workers[i].events) < 1 {
					return
				}
				timer := time.NewTimer(now.Sub(c.workers[i].events[0].At))
				<-timer.C
				c.workers[i].events[0].Job.Run()
			}
		}()
	}
}
