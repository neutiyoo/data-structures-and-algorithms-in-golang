package tree

import (
	"reflect"
	"testing"
)

func TestSerializeLevelorder(t *testing.T) {
	root := &Node{Key: 1}
	root.Left = &Node{Key: 2}
	root.Right = &Node{Key: 3}
	root.Left.Left = &Node{Key: 4}
	root.Left.Right = &Node{Key: 5}

	want := "1,2,3,4,5"

	result := Serialize(root)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}
}
func TestDeserializeLevelorder(t *testing.T) {
	x := "1,2,3,4,5"
	root := Deserialize(x)
	want := []int{1, 2, 3, 4, 5}

	result := LevelorderIterative(root)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}
}

func TestLevelorderIterative(t *testing.T) {
	x := "1,2,3,4,5"
	root := Deserialize(x)
	want := []int{1, 2, 3, 4, 5}

	result := LevelorderIterative(root)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}
}

func TestInorderIterative(t *testing.T) {
	result := InorderIterative(nil)
	if len(result) != 0 {
		t.Errorf("%#v; want %#v", len(result), 0)
	}

	x := "5,2,7,1,3,6,8"
	root := Deserialize(x)
	want := []int{1, 2, 3, 5, 6, 7, 8}

	result = InorderIterative(root)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}
}

func TestInorderRecursive(t *testing.T) {
	result := InorderRecursive(nil)
	if len(result) != 0 {
		t.Errorf("%#v; want %#v", len(result), 0)
	}

	x := "5,2,7,1,3,6,8"
	root := Deserialize(x)
	want := []int{1, 2, 3, 5, 6, 7, 8}

	result = InorderRecursive(root)

	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}
}

func TestSerializePreoder(t *testing.T) {
	root := &Node{Key: 1}
	root.Left = &Node{Key: 2}
	root.Right = &Node{Key: 3}
	root.Left.Left = &Node{Key: 4}
	root.Left.Right = &Node{Key: 5}
	root.Right.Left = &Node{Key: 6}
	root.Right.Right = &Node{Key: 7}

	want := "1,2,4,nil,nil,5,nil,nil,3,6,nil,nil,7"

	result := SerializePreoder(root)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}
}
func TestDeserializePreorder(t *testing.T) {
	x := "1,2,4,nil,nil,5,nil,nil,3,6,nil,nil,7"
	root := DeserializePreorder(x)
	want := []int{1, 2, 4, 5, 3, 6, 7}

	result := PreorderRecursive(root)
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}
}

func TestPreorderRecursive(t *testing.T) {
	result := PreorderRecursive(nil)
	if len(result) != 0 {
		t.Errorf("%#v; want %#v", len(result), 0)
	}

	x := "1,2,4,nil,nil,5,nil,nil,3,6,nil,nil,7"
	root := DeserializePreorder(x)

	result = PreorderRecursive(root)

	want := []int{1, 2, 4, 5, 3, 6, 7}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}

}

func TestPostorderRecursive(t *testing.T) {
	result := PostorderRecursive(nil)
	if len(result) != 0 {
		t.Errorf("%#v; want %#v", len(result), 0)
	}

	root := &Node{Key: 1}
	root.Left = &Node{Key: 2}
	root.Right = &Node{Key: 3}
	root.Left.Left = &Node{Key: 4}
	root.Left.Right = &Node{Key: 5}
	root.Right.Left = &Node{Key: 6}
	root.Right.Right = &Node{Key: 7}

	result = PostorderRecursive(root)

	want := []int{4, 5, 2, 6, 7, 3, 1}
	if !reflect.DeepEqual(result, want) {
		t.Errorf("%#v; want %#v", result, want)
	}
}

func TestGetHeight(t *testing.T) {
	root := &Node{Key: 5}

	h := GetHeight(root)
	if h != 0 {
		t.Errorf("%#v; want %#v", h, 0)
	}

	root.Left = &Node{Key: 2}

	h = GetHeight(root)
	if h != 1 {
		t.Errorf("%#v; want %#v", h, 1)
	}

	root.Left.Left = &Node{Key: 1}
	root.Left.Right = &Node{Key: 3}

	h = GetHeight(root)
	if h != 2 {
		t.Errorf("%#v; want %#v", h, 2)
	}

	root.Right = &Node{Key: 7}
	root.Right.Left = &Node{Key: 6}
	root.Right.Right = &Node{Key: 8}
	root.Right.Right.Right = &Node{Key: 9}

	h = GetHeight(root)
	if h != 3 {
		t.Errorf("%#v; want %#v", h, 3)
	}
}
