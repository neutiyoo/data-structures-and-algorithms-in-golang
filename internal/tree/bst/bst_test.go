package bst

import (
	"algorithms/internal/tree"
	"reflect"
	"testing"
)

func TestFromSortedArray(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5, 6, 7}
	root := FromSortedArray(sorted)
	for i, v := range tree.InorderIterative(root) {
		if sorted[i] != v {
			t.Errorf("%#v; want %#v", v, sorted[i])
		}
	}
}

func TestIsValid(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5, 6, 7}
	root := FromSortedArray(sorted)
	result := IsValid(root)
	if !result {
		t.Errorf("%#v; want %#v", result, true)
	}

	unsorted := []int{3, 2, 1, 7, 5, 6, 4}
	root = FromSortedArray(unsorted)
	result = IsValid(root)
	if result {
		t.Errorf("%#v; want %#v", result, false)
	}

	unsorted = []int{3, 2, 1, 1, 5, 6, 4}
	root = FromSortedArray(unsorted)

	result = IsValid(root)
	if result {
		t.Errorf("%#v; want %#v", result, false)
	}
}

func TestInsert(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 9}
	root := &tree.Node{Key: 6}
	for i, v := range nums {
		Insert(root, v, i)
	}
	want := []int{1, 3, 5, 6, 7, 9}
	result := tree.InorderRecursive(root)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, true)
	}
}

func TestSearch(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9}
	root := &tree.Node{Key: 4}
	for i, v := range nums {
		Insert(root, v, i)
	}
	result := Search(root, 7)
	if result != 3 {
		t.Errorf("%#v; want %#v", result, 3)
	}
}

func TestMin(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5, 6, 7}

	root := FromSortedArray(sorted)
	v := Min(root)
	if v != 1 {
		t.Errorf("%#v; want %#v", v, 1)
	}

	v = Min(nil)
	if v != nil {
		t.Errorf("%#v; want %#v", v, nil)
	}
}
func TestMax(t *testing.T) {
	sorted := []int{1, 2, 3, 4, 5, 6, 7}

	root := FromSortedArray(sorted)
	v := Max(root)
	if v != 7 {
		t.Errorf("%#v; want %#v", v, 7)
	}

	v = Max(nil)
	if v != nil {
		t.Errorf("%#v; want %#v", v, nil)
	}
}
