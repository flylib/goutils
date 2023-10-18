package grid

import (
	"sync"
	"sync/atomic"
)

// 二维网格
type Group[T any] struct {
	size     int
	elements []atomic.Pointer[T]
}

func NewGroup[T any](size int) *Group[T] {
	return &Group[T]{
		size:     size,
		elements: make([]atomic.Pointer[T], size),
	}
}

func (g *Group[T]) Load(i int) (*T, bool) {
	if i >= g.size {
		return nil, false
	}
	v := g.elements[i].Load()
	return v, v == nil
}

func (g *Group[T]) Set(i int, v *T) {
	if i >= g.size {
		return
	}
	g.elements[i].Store(v)
}

func (g *Group[T]) Del(i int) {
	if i >= g.size {
		return
	}
	g.elements[i].Store(nil)
}

type Grid[T any] struct {
	groupLength int
	size        int
	groups      Group[T]
	sync.Mutex
}

func NewGrid[T any](groupLength int) {

}
