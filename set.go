package goutils

// 定义泛型 Set
type Set[T comparable] map[T]struct{}

// 创建一个新 Set
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// 添加元素
func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

// 删除元素
func (s Set[T]) Remove(value T) {
	delete(s, value)
}

// 判断是否包含元素
func (s Set[T]) Contains(value T) bool {
	_, ok := s[value]
	return ok
}

// 获取所有元素
func (s Set[T]) Elements() []T {
	result := make([]T, 0, len(s))
	for k := range s {
		result = append(result, k)
	}
	return result
}
