package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := New()
	words := []string{
		"abc", "cat", "cats",
	}
	for _, word := range words {
		trie.Insert(word)
	}

	testCases := []struct {
		word string
		want bool
	}{
		{"abc", true},
		{"cat", true},
		{"cats", true},

		{"c", false},
		{"catss", false},
		{"d", false},
	}

	for _, testCase := range testCases {
		result := trie.Search(testCase.word)
		if result != testCase.want {
			t.Errorf("%#v; want %#v", result, testCase.want)
		}
	}

	prefixes := []string{
		"ab", "ca",
	}
	for _, prefix := range prefixes {
		result := trie.StartsWith(prefix)
		if !result {
			t.Errorf("%#v; want %#v", result, true)
		}
	}
}
