package core

type Stack[T any] struct {
	data []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: make([]T, 0, 16)}
}

func (s *Stack[T]) Push(data T) {
	s.data = append(s.data, data)
}

func (s *Stack[T]) Pop() (T, bool) {
	data := s.data
	if len(data) == 0 {
		var zero T
		return zero, false
	}

	i := len(data) - 1
	val := data[i]
	s.data = data[:i]
	return val, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}
