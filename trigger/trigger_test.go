package trigger

import (
	"fmt"
	"testing"
	"time"
)

type EveryMintTask struct {
}

var i int64

func (EveryMintTask) Run() {
	i++
	fmt.Println(time.Now(), "--EveryMintTask--", i)
}

func TestNewTrigger(t *testing.T) {
	trigger := NewTrigger()
	now := time.Now()
	for i := 0; i < 10; i++ {
		now = now.Add(time.Millisecond)
		task := Task{time: now, Job: EveryMintTask{}}
		trigger.AddTask(task)
	}
	trigger.Start()
}
