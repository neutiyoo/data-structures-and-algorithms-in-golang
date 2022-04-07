package heap

import (
	"testing"
)

func TestMinHeap(t *testing.T) {
	testCases := []struct {
		insertTests []struct {
			val     int
			lenWant int
		}
		extractTests []struct {
			keyWant interface{}
			lenWant int
		}
	}{
		{
			insertTests: []struct {
				val     int
				lenWant int
			}{
				{5, 1},
				{3, 2},
				{1, 3},
				{4, 4},
				{2, 5},
			},
			extractTests: []struct {
				keyWant interface{}
				lenWant int
			}{
				{1, 4},
				{2, 3},
				{3, 2},
				{4, 1},
				{5, 0},
				{nil, 0},
			},
		},
	}

	for _, testCase := range testCases {
		mh := Heap{}
		for _, insertTest := range testCase.insertTests {
			mh.Insert(insertTest.val, insertTest.val)
			if mh.Len() != insertTest.lenWant {
				t.Errorf("%#v %#v", mh.Len(), insertTest.lenWant)
			}
		}

		for _, extractTest := range testCase.extractTests {
			el := mh.Extract()
			if extractTest.keyWant == nil {
				if el != nil {
					t.Errorf("%#v %#v", el, extractTest.keyWant)
				}
			} else {
				if el.Key != extractTest.keyWant {
					t.Errorf("%#v %#v", el.Key, extractTest.keyWant)
				}
			}
		}
	}
}
