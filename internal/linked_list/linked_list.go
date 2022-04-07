package linkedlist

type Node struct {
	Val  interface{}
	Next *Node
}

type SinglyLinkedList struct {
	head *Node
	size int
}

func (l *SinglyLinkedList) Len() int {
	return l.size
}

func (l *SinglyLinkedList) Insert(value interface{}) {
	newNode := Node{Val: value}
	l.size++
	if l.head == nil {
		l.head = &newNode
	}
	newNode.Next = l.head
	l.head = &newNode
}

func (l *SinglyLinkedList) Delete(i int) interface{} {
	if i < 0 || i > l.size-1 {
		return nil
	}
	l.size--

	// Remove first node
	if i == 0 {
		value := l.head.Val
		l.head = l.head.Next
		return value
	}

	// Navigate to the i-1 node
	cur := l.head
	j := 0
	for j < i-1 {
		j++
		cur = cur.Next
	}

	value := cur.Next.Val
	cur.Next = cur.Next.Next
	return value
}

func FromArray(a []interface{}) *Node {
	if len(a) == 0 {
		return nil
	}
	head := &Node{Val: a[0]}
	cur := head
	for i := 1; i < len(a); i++ {
		newNode := &Node{Val: a[i]}
		cur.Next = newNode
		cur = cur.Next
	}
	return head
}

func ToArray(h *Node) []interface{} {
	if h == nil {
		return []interface{}{}
	}

	result := []interface{}{}
	cur := h
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}
