package trigger

import "sync"

// item is an element of a linked list.
type item struct {
	next *item
	list *List
	task Task
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root *item // sentinel list element, only &root, root.prev, and root.next are used
	len  int   // current list length excluding (this) sentinel element
	*sync.Mutex
}

func newList() List {
	return List{
		root:  &item{},
		len:   0,
		Mutex: &sync.Mutex{},
	}
}

// Len returns the number of elements of list l.
// The complexity is O(1).
func (l *List) Len() int { return l.len }

// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e *item) *item {
	l.Lock()
	defer l.Unlock()
	next := l.root.next
	prev := l.root
	for {
		if next == nil {
			next = e
			prev.next = next
			break
		}
		//时间小的靠前
		if e.task.time.Before(next.task.time) {
			tmp := l.root.next
			l.root.next = e
			e.next = tmp
			break
		}
		prev = next
		next = next.next
	}
	l.len++
	return e
}

func (l *List) pop() *item {
	l.Lock()
	defer l.Unlock()
	e := l.root.next
	if l.len > 0 {
		l.root.next = e.next
		e.next = nil
		l.len--
	}
	return e
}
func (l *List) Front() *item {
	return l.root.next
}
