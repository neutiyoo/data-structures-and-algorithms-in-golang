package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := New()
	if s.Len() != 0 {
		t.Errorf("%#v; want %#v", s.Len(), 0)
	}
	if s.Peek() != nil {
		t.Errorf("%#v; want %#v", s.Peek(), nil)
	}

	testCasesPush := []struct {
		val     int
		lenWant int
	}{
		{1, 1},
		{3, 2},
		{5, 3},
	}
	for _, testCase := range testCasesPush {

		s.Push(testCase.val)
		result := s.Len()
		if result != testCase.lenWant {
			t.Errorf("%#v; want %#v", result, testCase.lenWant)
		}
		if s.Peek() != testCase.val {
			t.Errorf("%#v; want %#v", result, testCase.val)
		}
	}

	testCasesPop := []struct {
		valWant interface{}
		lenWant int
	}{
		{5, 2},
		{3, 1},
		{1, 0},
		{nil, 0},
	}
	for _, testCase := range testCasesPop {
		el := s.Pop()
		if el != testCase.valWant {
			t.Errorf("%#v; want %#v", el, testCase.valWant)
		}
		l := s.Len()
		if l != testCase.lenWant {
			t.Errorf("%#v; want %#v", l, testCase.lenWant)
		}
	}
}
