package stack

import (
	linkedlist "algorithms/internal/linked_list"
)

type Stack struct {
	size int
	head *linkedlist.Node
}

func New() Stack {
	return Stack{}
}

func (s Stack) Len() int {
	return s.size
}

// Time complexity: O(1)
func (s Stack) Peek() interface{} {
	if s.head == nil {
		return nil
	}
	return s.head.Val
}

// Time complexity: O(1)
func (s *Stack) Push(v interface{}) {
	newNode := linkedlist.Node{
		Val:  v,
		Next: nil,
	}
	s.size++

	if s.head == nil {
		s.head = &newNode
		return
	}

	newNode.Next = s.head
	s.head = &newNode

}

// Time complexity: O(1)
func (s *Stack) Pop() interface{} {
	if s.head == nil {
		return nil
	}

	s.size--

	v := s.head.Val
	s.head = s.head.Next
	return v
}
