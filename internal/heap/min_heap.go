package heap

import (
	"algorithms/internal/tree"
)

func getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}
func getLeftChildIndex(parentIndex int) int {
	return parentIndex*2 + 1
}
func getRightChildIndex(parentIndex int) int {
	return parentIndex*2 + 2
}

type Heap struct {
	arr  []tree.Node
	size int
}

func (h *Heap) Len() int {
	return h.size
}

// Time complexity: O(log n)
func (h *Heap) Insert(key int, value interface{}) {
	newNode := tree.Node{
		Key: key,
		Val: value,
	}
	h.arr = append(h.arr, newNode)
	h.size++
	h.heapifyUp()
}

func (h *Heap) heapifyUp() {
	childIndex := len(h.arr) - 1
	for getParentIndex(childIndex) > -1 {
		parentIndex := getParentIndex(childIndex)
		if h.arr[parentIndex].Key <= h.arr[childIndex].Key {
			break
		}
		h.arr[parentIndex], h.arr[childIndex] = h.arr[childIndex], h.arr[parentIndex]
		childIndex = parentIndex
	}
}

// Time complexity: O(log n)
func (h *Heap) Extract() *tree.Node {
	if len(h.arr) == 0 {
		return nil
	}
	node := h.arr[0]
	h.arr[0], h.arr[len(h.arr)-1] = h.arr[len(h.arr)-1], h.arr[0]
	h.arr = h.arr[:len(h.arr)-1]
	h.size--
	h.heapifyDown()
	return &node
}

func (h *Heap) heapifyDown() {
	parentIndex := 0
	for getLeftChildIndex(parentIndex) < len(h.arr) {
		left := getLeftChildIndex(parentIndex)
		right := getRightChildIndex((parentIndex))
		smallChildIndex := left

		if right < len(h.arr) && h.arr[right].Key < h.arr[left].Key {
			smallChildIndex = right
		}

		if h.arr[parentIndex].Key <= h.arr[smallChildIndex].Key {
			break
		}
		h.arr[parentIndex], h.arr[smallChildIndex] = h.arr[smallChildIndex], h.arr[parentIndex]
		parentIndex = smallChildIndex
	}
}
