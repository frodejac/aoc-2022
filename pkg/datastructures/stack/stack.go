package stack

type Stack[T any] struct {
	values []T
}

func New[T any](values []T) *Stack[T] {
	return &Stack[T]{values: values}
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) PushLeft(value T) {
	s.values = append([]T{value}, s.values...)
}

func (s *Stack[T]) Pop() T {
	if len(s.values) == 0 {
		panic("stack is empty")
	}
	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return value
}

func (s *Stack[T]) PopLeft() T {
	if len(s.values) == 0 {
		panic("stack is empty")
	}
	value := s.values[0]
	s.values = s.values[1:]
	return value
}

func (s *Stack[T]) PopN(count int) []T {
	if len(s.values) < count {
		panic("not enough values")
	}
	values := s.values[len(s.values)-count:]
	s.values = s.values[:len(s.values)-count]
	return values
}

func (s *Stack[T]) PushN(values []T) {
	s.values = append(s.values, values...)
}

func (s *Stack[T]) Peek() T {
	return s.values[len(s.values)-1]
}

func (s *Stack[T]) Reverse() {
	for i := 0; i < len(s.values)/2; i++ {
		s.values[i], s.values[len(s.values)-1-i] = s.values[len(s.values)-1-i], s.values[i]
	}
}
