package hashtable

import (
	"testing"
)

func TestHashTable(t *testing.T) {

	h := New(3)

	testCases := []struct {
		key string
		val int
	}{
		{"sarang", 0},
		{"ahri", 1},
		{"asher", 2},
		{"ariel", 4},
	}
	for _, testCase := range testCases {
		// Add new nodes
		h.Set(testCase.key, testCase.val)
		result := h.Get(testCase.key)
		if result != testCase.val {
			t.Errorf("%#v; want %#v", result, testCase.val)
		}
	}

	// Update existing node
	h.Set("ahri", 8)
	result := h.Get("ahri")
	if result != 8 {
		t.Errorf("%#v; want %#v", result, 8)
	}

	result = h.Get("cassy")
	if result != nil {
		t.Errorf("%#v; want %#v", result, nil)
	}
	result = h.Get("bob")
	if result != nil {
		t.Errorf("%#v; want %#v", result, nil)
	}

	// Delete nodes
	result = h.Delete("bob")
	if result != false {
		t.Errorf("%#v; want %#v", result, false)
	}

	keys := []string{
		"sarang",
		"asher",
		"ariel",
		"ahri",
	}
	for _, key := range keys {
		result = h.Delete(key)
		if result != true {
			t.Errorf("%#v; want %#v", result, true)
		}
		result = h.Delete(key)
		if result != false {
			t.Errorf("%#v; want %#v", result, false)
		}
		result = h.Get(key)
		if result != nil {
			t.Errorf("%#v; want %#v", result, nil)
		}
	}
}
