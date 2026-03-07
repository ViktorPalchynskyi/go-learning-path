package main

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int      { return len(s.items) }
func (s *Stack[T]) IsEmpty() bool { return len(s.items) == 0 }

func main() {
	// Stack[int]
	var si Stack[int]
	si.Push(1)
	si.Push(2)
	si.Push(3)
	fmt.Printf("int stack len=%d, isEmpty=%v\n", si.Len(), si.IsEmpty())

	if top, ok := si.Peek(); ok {
		fmt.Printf("peek: %d\n", top)
	}
	for {
		v, ok := si.Pop()
		if !ok {
			break
		}
		fmt.Printf("pop: %d\n", v)
	}
	fmt.Printf("int stack after drain: isEmpty=%v\n", si.IsEmpty())

	// Pop on empty stack
	_, ok := si.Pop()
	fmt.Printf("pop on empty: ok=%v\n", ok)

	// Stack[string]
	var ss Stack[string]
	ss.Push("hello")
	ss.Push("world")
	fmt.Printf("\nstring stack len=%d\n", ss.Len())
	if top, ok := ss.Peek(); ok {
		fmt.Printf("peek: %s\n", top)
	}
	if v, ok := ss.Pop(); ok {
		fmt.Printf("pop: %s\n", v)
	}
}
