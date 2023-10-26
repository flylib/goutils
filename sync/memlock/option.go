package memlock

import (
	"time"
)

type Option func(lock *RWLock)

// Default no clean up cache
func WithCleanUpInterval(d time.Duration) Option {
	return func(lock *RWLock) {
		lock.cleanUpInterval = d
	}
}

// Default 30s will expired
func WithMaxLockDuration(d time.Duration) Option {
	return func(lock *RWLock) {
		lock.lockDuration = d
	}
}
