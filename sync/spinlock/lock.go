package spinlock

import (
	"runtime"
	"sync/atomic"
)

// A Locker must not be copied after first use.
type Locker struct {
	//_       sync.Mutex // for copy protection compiler warning
	locking atomic.Bool
}

// Lock locks l.
// If the lock is already in use, the calling goroutine
// blocks until the locker is available.
func (l *Locker) Lock() {
	for l.locking.Swap(true) {
		runtime.Gosched()
	}
}

// Unlock unlocks l.
func (l *Locker) Unlock() {
	l.locking.Store(false)
}
