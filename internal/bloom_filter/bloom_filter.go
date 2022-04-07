package bloomfilter

import (
	"hash/fnv"
)

type BloomFilter struct {
	bucket []uint64
	size   uint64
}

func New(size uint64) BloomFilter {
	b := BloomFilter{}
	b.bucket = make([]uint64, size)
	b.size = size
	return b
}

func (b *BloomFilter) hash(s string) []uint64 {
	h1 := fnv.New64()
	h1.Write([]byte(s))
	x := h1.Sum64()
	x %= b.size

	h2 := fnv.New64a()
	h2.Write([]byte(s))
	y := h2.Sum64()
	y %= b.size

	return []uint64{x, y}
}

func (b *BloomFilter) Set(s string) {
	for _, v := range b.hash(s) {
		b.bucket[v] = 1
	}
}

func (b *BloomFilter) ProbablyHas(s string) bool {
	for _, v := range b.hash(s) {
		if b.bucket[v] != 1 {
			return false
		}
	}
	return true
}
