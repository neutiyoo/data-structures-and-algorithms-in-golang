package queue

import (
	linkedlist "algorithms/internal/linked_list"
)

type Queue struct {
	size int
	head *linkedlist.Node
	tail *linkedlist.Node
}

func New() Queue {
	return Queue{}
}

func (s Queue) Len() int {
	return s.size
}

// Time complexity: O(1)
func (s Queue) Peek() interface{} {
	if s.head == nil {
		return nil
	}
	return s.head.Val
}

// Time complexity: O(1)
func (s *Queue) Enqueue(v interface{}) {
	newNode := linkedlist.Node{
		Val:  v,
		Next: nil,
	}
	s.size++

	if s.head == nil {
		s.head = &newNode
		s.tail = &newNode
		return
	}

	s.tail.Next = &newNode
	s.tail = &newNode
}

// Time complexity: O(1)
func (s *Queue) Dequeue() interface{} {
	if s.head == nil {
		return nil
	}

	s.size--

	v := s.head.Val
	if s.head.Next == nil {
		s.head = nil
		s.tail = nil
		return v
	}

	s.head = s.head.Next
	return v
}
