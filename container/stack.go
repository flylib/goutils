package container

import (
	"sync"
)

type Item[T any] struct {
	next *Item[T]
	v    T
}

// 栈
type Stack[T any] struct {
	root *Item[T]
	sync.Mutex
}

func (p *Stack[T]) init() *Stack[T] {
	p.root = new(Item[T])
	p.root.next = nil
	return p
}

func NewStack[T any]() *Stack[T] { return new(Stack[T]).init() }

// 压入
func (p *Stack[T]) Push(v T) {
	e := &Item[T]{v: v}
	p.Lock()
	e.next = p.root.next
	p.root.next = e
	p.Unlock()
}

// 弹出
func (p *Stack[T]) Pop() T {
	if p.root.next == nil {
		var t T
		return t
	}
	p.Lock()
	e := p.root.next
	p.root.next = e.next
	p.Unlock()
	return e.v
}
