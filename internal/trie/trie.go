package trie

type Node struct {
	Children map[byte]*Node
	IsEnd    bool
}

type Trie struct {
	root *Node
}

func New() Trie {
	t := Trie{}
	t.root = &Node{
		IsEnd:    false,
		Children: make(map[byte]*Node),
	}
	return t
}

func (t *Trie) Root() *Node {
	return t.root
}

// Time complexity: O(m)
// Space complexity: O(m)
// - m is the length of a search string
func (t *Trie) Insert(word string) {
	cur := t.root
	for i := 0; i < len(word); i++ {
		if _, ok := cur.Children[word[i]]; !ok {
			cur.Children[word[i]] = &Node{
				IsEnd:    false,
				Children: make(map[byte]*Node),
			}
		}

		cur = cur.Children[word[i]]
	}

	cur.IsEnd = true
}

// Time complexity: O(m)
// Space complexity: O(1)
// - m is the length of a search string
func (t *Trie) Search(word string) bool {
	cur := t.root
	for i := 0; i < len(word); i++ {
		if v := cur.Children[word[i]]; v == nil {
			return false
		}
		cur = cur.Children[word[i]]
	}

	return cur.IsEnd
}

// Time complexity: O(m)
// Space complexity: O(1)
// - m is the length of a search string
func (t *Trie) StartsWith(prefix string) bool {
	cur := t.root
	for i := 0; i < len(prefix); i++ {
		if v := cur.Children[prefix[i]]; v == nil {
			return false
		}
		cur = cur.Children[prefix[i]]
	}

	return true
}
