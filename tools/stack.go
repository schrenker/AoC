package tools

type Stack[T any] struct {
	array []T
	top   int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{top: 0, array: make([]T, 0)}
}

func (s *Stack[T]) Push(elem T) *Stack[T] {
	s.array = append(s.array, elem)
	s.top++
	return s
}

func (s *Stack[T]) Pop() T {
	elem := s.array[len(s.array)-1]
	s.array = s.array[0 : s.top-1]
	s.top--
	return elem
}

func (s *Stack[T]) Peek() T {
	return s.array[s.Length()-1]
}

func (s *Stack[T]) Length() int {
	return len(s.array)
}

func (s *Stack[T]) List() []T {
	return s.array
}

func (s *Stack[T]) Reverse() *Stack[T] {
	s.array = Reverse(s.array)
	return s
}
