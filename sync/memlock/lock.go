package memlock

import (
	"runtime"
	"sync"
	"time"
)

type RWLock struct {
	cleanUpInterval time.Duration
	lockDuration    time.Duration
	story           map[string]int64
	mutex           sync.Mutex
}

func NewRWLock(options ...Option) *RWLock {
	l := RWLock{
		lockDuration: time.Second * 30,
		story:        map[string]int64{},
	}
	for _, f := range options {
		f(&l)
	}
	if l.cleanUpInterval > 0 {
		go func() {
			ticker := time.NewTicker(l.cleanUpInterval)
			for range ticker.C {
				now := time.Now().UnixNano()
				l.mutex.Lock()
				for k, v := range l.story {
					if v < now {
						delete(l.story, k)
					}
				}
				l.mutex.Unlock()
			}
		}()
	}

	return &l
}

func (r *RWLock) Lock(name string) {
	for {
		now := time.Now()
		r.mutex.Lock()
		at, ok := r.story[name]
		if ok && now.UnixNano() < at {
			r.mutex.Unlock()
			runtime.Gosched()
			continue
		}
		r.story[name] = now.Add(r.lockDuration).UnixNano()
		r.mutex.Unlock()
		return
	}
}

func (r *RWLock) UnLock(name string) {
	r.mutex.Lock()
	r.story[name] = 0
	r.mutex.Unlock()
}
