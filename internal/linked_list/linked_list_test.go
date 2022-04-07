package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := SinglyLinkedList{}

	v := l.Delete(0)
	if v != nil {
		t.Errorf("want %v", nil)
	}
	if l.Len() != 0 {
		t.Errorf("s.Len() = %#v; want %#v", l.Len(), 0)
	}

	testCasesInsert := []struct {
		val  int
		want int
	}{
		{1, 1},
		{3, 2},
		{5, 3},
		{2, 4},
		{4, 5},
	}
	for _, testCase := range testCasesInsert {

		l.Insert(testCase.val)
		length := l.Len()
		if length != testCase.want {
			t.Errorf("%#v; want %#v", length, testCase.want)
		}
	}

	cur := l.head
	for _, v := range []int{4, 2, 5, 3, 1} {
		if cur.Val.(int) != v {
			t.Errorf("%#v; want %#v", cur.Val.(int), v)
		}
		cur = cur.Next
	}

	testCasesDelete := []struct {
		index     int
		valWant   int
		lenWant   int
		sliceWant []int
	}{
		{0, 4, 4, []int{2, 5, 3, 1}},
		{3, 1, 3, []int{2, 5, 3}},
		{1, 5, 2, []int{2, 3}},
	}
	for _, testCase := range testCasesDelete {
		result := l.Delete(testCase.index)
		if result != testCase.valWant {
			t.Errorf("%#v; want %#v", result, testCase.valWant)
		}
		length := l.Len()
		if length != testCase.lenWant {
			t.Errorf("%#v; want %#v", length, testCase.lenWant)
		}
		cur = l.head
		for _, val := range testCase.sliceWant {
			if cur.Val.(int) != val {
				t.Errorf("%#v; want %#v", cur.Val.(int), val)
			}
			cur = cur.Next
		}
	}
}

func TestFromArray(t *testing.T) {
	x := []interface{}{}
	head := FromArray(x)
	if head != nil {
		t.Errorf("%#v; want %#v", head, nil)
	}

	x = []interface{}{1, 2, 3, 4, 5}
	head = FromArray(x)

	cur := head
	for _, want := range []interface{}{1, 2, 3, 4, 5} {
		v := cur.Val.(int)
		if !reflect.DeepEqual(cur.Val, v) {
			t.Errorf("%#v; want %#v", v, want)
		}
		cur = cur.Next
	}
	if cur != nil {
		t.Errorf("%#v; want %#v", cur, nil)
	}
}

func TestToArray(t *testing.T) {
	want := []interface{}{}
	result := ToArray(nil)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}

	want = []interface{}{1, 2, 3, 4, 5}
	head := &Node{}
	cur := head
	for _, v := range want {
		newNode := &Node{Val: v}
		cur.Next = newNode
		cur = cur.Next
	}

	result = ToArray(head.Next)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}

}
