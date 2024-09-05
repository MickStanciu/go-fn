package stack

import "fmt"

type Stack[T any] struct {
	content []T
}

// NewStack - creates a new empty stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push - pushes an element to the stack
func (s *Stack[T]) Push(item T) {
	s.content = append(s.content, item)
}

// Pop - pops an element from the stack
func (s *Stack[T]) Pop() (T, error) {
	if len(s.content) == 0 {
		return *new(T), fmt.Errorf("stack is empty")
	}
	lastIdx := len(s.content) - 1
	lastItem := s.content[lastIdx]
	s.content = s.content[:lastIdx]
	return lastItem, nil
}

// GetLength - returns the size of the stack
func (s *Stack[T]) GetLength() int {
	return len(s.content)
}
