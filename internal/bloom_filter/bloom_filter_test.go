package bloomfilter

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	b := New(100)
	b.Set("ahri")
	result := b.ProbablyHas("ahri")
	if !result {
		t.Errorf("%#v; want %#v", result, true)
	}

	result = b.ProbablyHas("bob")
	if result {
		t.Errorf("%#v; want %#v", result, false)
	}

	b.Set("bob")
	result = b.ProbablyHas("bob")
	if !result {
		t.Errorf("%#v; want %#v", result, true)
	}
}
