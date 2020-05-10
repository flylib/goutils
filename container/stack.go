package util

import (
	"sync"
)

type e struct {
	next    *e
	content interface{}
}

//栈
type Stack struct {
	root *e
	sync.Mutex
}

func (p *Stack) init() *Stack {
	p.root = new(e)
	p.root.next = nil
	return p
}

func NewStack() *Stack { return new(Stack).init() }

//压入
func (p *Stack) Push(v interface{}) {
	e := &e{content: v}
	p.Lock()
	e.next = p.root.next
	p.root.next = e
	p.Unlock()
}

//弹出
func (p *Stack) Pop() interface{} {
	if p.root.next == nil {
		return nil
	}
	p.Lock()
	e := p.root.next
	p.root.next = e.next
	p.Unlock()
	return e.content
}
