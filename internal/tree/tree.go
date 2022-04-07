package tree

import (
	"container/list"
	"strconv"
	"strings"
)

type Node struct {
	Key   int
	Val   interface{}
	Left  *Node
	Right *Node
}

func LevelorderIterative(root *Node) []int {
	result := []int{}
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		el := queue.Remove(queue.Front())
		node := el.(*Node)

		if node == nil {
			continue
		}

		result = append(result, node.Key)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
	return result
}

func InorderIterative(node *Node) []int {
	result := []int{}
	if node == nil {
		return result
	}

	stack := list.New()
	cur := node

	for stack.Len() > 0 || cur != nil {
		if cur != nil {
			stack.PushFront(cur)
			cur = cur.Left
			continue
		}

		el := stack.Remove(stack.Front())
		node := el.(*Node)
		result = append(result, node.Key)
		cur = node.Right
	}
	return result
}

func PreorderRecursive(root *Node) []int {
	result := []int{}

	var preorder func(*Node, *[]int)
	preorder = func(node *Node, res *[]int) {
		if node == nil {
			return
		}
		*res = append(*res, node.Key)
		preorder(node.Left, res)
		preorder(node.Right, res)
	}

	preorder(root, &result)
	return result
}

func InorderRecursive(root *Node) []int {
	result := []int{}

	var inorder func(*Node, *[]int)
	inorder = func(node *Node, res *[]int) {
		if node == nil {
			return
		}
		inorder(node.Left, res)
		*res = append(*res, node.Key)
		inorder(node.Right, res)
	}

	inorder(root, &result)
	return result
}

func PostorderRecursive(root *Node) []int {
	result := []int{}

	var postorder func(*Node, *[]int)
	postorder = func(node *Node, res *[]int) {
		if node == nil {
			return
		}
		postorder(node.Left, res)
		postorder(node.Right, res)
		*res = append(*res, node.Key)
	}

	postorder(root, &result)
	return result
}

// height: number of edges between root node and furthest leaf
func GetHeight(node *Node) int {
	if node == nil {
		return -1
	}
	l := GetHeight(node.Left)
	r := GetHeight(node.Right)
	if l < r {
		return r + 1
	}
	return l + 1
}

// Level order
func Serialize(root *Node) string {
	if root == nil {
		return ""
	}
	queue := list.New()
	queue.PushBack(root)

	strKeys := []string{}
	for queue.Len() > 0 {
		el := queue.Remove(queue.Front())
		node := el.(*Node)
		if node == nil {
			strKeys = append(strKeys, "nil")
			continue
		}
		strKeys = append(strKeys, strconv.Itoa(node.Key))
		queue.PushBack(node.Left)
		queue.PushBack(node.Right)
	}

	for i := len(strKeys) - 1; i > -1; i-- {
		if strKeys[i] == "nil" {
			strKeys = strKeys[:i]
		} else {
			break
		}
	}
	return strings.Join(strKeys, ",")
}

// Level order
func Deserialize(s string) *Node {
	strKeys := strings.Split(s, ",")
	if len(s) == 0 {
		return nil
	}

	rootKey, err := strconv.Atoi(strKeys[0])
	if err != nil {
		return nil
	}
	root := &Node{Key: rootKey}

	queue := list.New()
	queue.PushBack(root)

	i := 1
	for i < len(strKeys) && queue.Len() > 0 {
		el := queue.Remove(queue.Front())
		node := el.(*Node)
		if strKeys[i] != "nil" {
			intKey, _ := strconv.Atoi(strKeys[i])
			newNode := &Node{Key: intKey}
			node.Left = newNode
			queue.PushBack(newNode)
		}
		i++

		if i >= len(strKeys) {
			break
		}
		if strKeys[i] != "nil" {
			intKey, _ := strconv.Atoi(strKeys[i])
			newNode := &Node{Key: intKey}
			node.Right = newNode
			queue.PushBack(newNode)
		}
		i++
	}

	return root
}

func SerializePreoder(root *Node) string {
	strKeys := []string{}
	var serialize func(*Node, *[]string)
	serialize = func(node *Node, keys *[]string) {
		if node == nil {
			*keys = append(*keys, "nil")
			return
		}
		*keys = append(*keys, strconv.Itoa(node.Key))
		serialize(node.Left, keys)
		serialize(node.Right, keys)
	}
	serialize(root, &strKeys)

	for i := len(strKeys) - 1; i > -1; i-- {
		if strKeys[i] == "nil" {
			strKeys = strKeys[:i]
		} else {
			break
		}
	}
	return strings.Join(strKeys, ",")
}

func DeserializePreorder(s string) *Node {
	strKeys := strings.Split(s, ",")

	var deserialize func(*[]string) *Node
	deserialize = func(keys *[]string) *Node {
		if len(*keys) == 0 {
			return nil
		}
		strKey := (*keys)[0]
		*keys = (*keys)[1:]

		if strKey == "nil" {
			return nil
		}

		intKey, _ := strconv.Atoi(strKey)

		node := &Node{Key: intKey}
		node.Left = deserialize(keys)
		node.Right = deserialize(keys)
		return node
	}

	root := deserialize(&strKeys)
	return root
}
