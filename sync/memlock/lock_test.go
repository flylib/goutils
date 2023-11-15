package memlock

import (
	"sync"
	"testing"
	"time"
)

// === RUN   TestLock
// --- PASS: TestLock (0.45s) =  10000
// lock_test.go:11: total =  100000
// --- PASS: TestLock (1.13s)
// === RUN   TestSLock
// lock_test.go:17: total =  100000
// --- PASS: TestSLock (1.80s)
// PASS
// ok      github.com/flylib/goutils/sync/memlock  3.531s
var lock = NewRWLock(WithMaxLockDuration(time.Second * 3))

func TestLock(t *testing.T) {
	t.Log(time.Now())
	lock.Lock("test")
	lock.UnLock("test")
	lock.Lock("test")
	t.Log(time.Now())
}

func TestLock1000(t *testing.T) {
	total := testLock("TestLock1000", 1000, lock)
	t.Log("total = ", total)
}
func TestLock10000(t *testing.T) {
	total := testLock("TestLock10000", 10000, lock)
	t.Log("total = ", total)
}

func TestLock100000(t *testing.T) {
	total := testLock("TestLock100000", 100000, lock)
	t.Log("total = ", total)
}

func TestLock1000000(t *testing.T) {
	total := testLock("TestLock1000000", 1000000, lock)
	t.Log("total = ", total)
}

func testLock(name string, n int, l *RWLock) int {
	w := sync.WaitGroup{}
	var total int
	for i := 0; i < n; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			l.Lock(name)
			total++
			l.UnLock(name)
		}()
	}
	w.Wait()
	return total
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
