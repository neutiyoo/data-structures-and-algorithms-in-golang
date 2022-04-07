package queue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := New()
	if q.Len() != 0 {
		t.Errorf("s.Len() = %#v; want %#v", q.Len(), 0)
	}
	if q.Peek() != nil {
		t.Errorf("s.Peek() = %#v; want %#v", q.Peek(), nil)
	}

	testCasesEnqueue := []struct {
		val     int
		lenWant int
		valWant int
	}{
		{1, 1, 1},
		{3, 2, 1},
		{5, 3, 1},
	}
	for _, testCase := range testCasesEnqueue {
		q.Enqueue(testCase.val)
		l := q.Len()
		if l != testCase.lenWant {
			t.Errorf("%#v; want %#v", l, testCase.lenWant)
		}
		p := q.Peek()
		if p != testCase.valWant {
			t.Errorf("%#v; want %#v", p, testCase.valWant)
		}
	}

	testCasesDeqeue := []struct {
		valWant interface{}
		lenWant int
	}{
		{1, 2},
		{3, 1},
		{5, 0},
		{nil, 0},
	}
	for _, testCase := range testCasesDeqeue {

		p := q.Peek()
		if p != testCase.valWant {
			t.Errorf("%#v; want %#v", p, testCase.valWant)
		}
		el := q.Dequeue()
		if testCase.valWant != el {
			t.Errorf("%#v; want %#v", el, testCase.valWant)
		}
		l := q.Len()
		if l != testCase.lenWant {
			t.Errorf("%#v; want %#v", l, testCase.lenWant)
		}
	}
}
