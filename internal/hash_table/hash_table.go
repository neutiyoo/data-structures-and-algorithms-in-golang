package hashtable

import (
	"hash/fnv"
)

type Node struct {
	Key   string
	Value interface{}
	Next  *Node
}

type HashTable struct {
	buckets []*Node
	size    uint64
}

func New(size uint64) HashTable {
	h := HashTable{}
	arr := make([]*Node, size)
	h.buckets = arr
	h.size = size
	return h
}

func (h *HashTable) hash(s string) uint64 {
	h1 := fnv.New64()
	h1.Write([]byte(s))
	x := h1.Sum64()
	x %= h.size
	return x
}

func (h *HashTable) Set(key string, value interface{}) {
	index := h.hash(key)
	newNode := &Node{Key: key, Value: value}
	// Init head
	if h.buckets[index] == nil {
		h.buckets[index] = newNode
		return
	}

	// Update value if key exists
	cur := h.buckets[index]
	for cur != nil {
		if cur.Key == key {
			cur.Value = value
			return
		}
		cur = cur.Next
	}

	// Add node from the head
	newNode.Next = h.buckets[index]
	h.buckets[index] = newNode
}

func (h *HashTable) Get(key string) interface{} {
	index := h.hash(key)
	cur := h.buckets[index]
	if cur == nil {
		return nil
	}
	for cur != nil {
		if cur.Key == key {
			return cur.Value
		}
		cur = cur.Next
	}
	return nil
}

func (h *HashTable) Delete(key string) bool {
	index := h.hash(key)
	head := h.buckets[index]

	if head == nil {
		return false
	}
	if head.Key == key {
		h.buckets[index] = head.Next
		return true
	}

	prev := head
	for prev.Next != nil {
		if prev.Next.Key == key {
			prev.Next = prev.Next.Next
			return true
		}
		prev = prev.Next
	}

	return false
}
