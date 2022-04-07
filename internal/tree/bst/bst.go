package bst

import (
	"algorithms/internal/tree"
	"sort"
)

func Insert(root *tree.Node, key int, value interface{}) {
	if Search(root, key) != nil {
		return
	}

	newNode := tree.Node{Key: key, Val: value}
	if key < root.Key {
		if root.Left == nil {
			root.Left = &newNode
		} else {
			Insert(root.Left, key, value)
		}
	} else {
		if root.Right == nil {
			root.Right = &newNode
		} else {
			Insert(root.Right, key, value)
		}
	}
}

func Search(root *tree.Node, key int) interface{} {
	if root.Key == key {
		return root.Val
	}
	if key < root.Key && root.Left != nil {
		return Search(root.Left, key)
	}
	if key > root.Key && root.Right != nil {
		return Search(root.Right, key)
	}
	return nil
}

func Min(root *tree.Node) interface{} {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root.Key
	}
	return Min(root.Left)
}

func Max(root *tree.Node) interface{} {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root.Key
	}
	return Max(root.Right)
}

func FromSortedArray(arr []int) *tree.Node {
	if len(arr) == 0 {
		return nil
	}
	mid := len(arr) / 2
	node := tree.Node{Key: arr[mid]}
	node.Left = FromSortedArray(arr[:mid])
	node.Right = FromSortedArray(arr[mid+1:])
	return &node
}

func IsValid(root *tree.Node) bool {
	inorder := tree.InorderIterative(root)
	keyMap := map[int]bool{}

	for _, v := range inorder {
		if !keyMap[v] {
			keyMap[v] = true
		} else {
			// All keys should be distinct
			return false
		}
	}

	sorted := make([]int, len(inorder))
	copy(sorted, inorder)
	sort.Ints(sorted)

	for i, v := range inorder {
		if sorted[i] != v {
			return false
		}
	}
	return true
}
