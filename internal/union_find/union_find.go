package unionfind

// | Operation                                    | Time Complexity |
// | -------------------------------------------- | :-------------: |
// | UNION+FIND (union by rank, path compression) |   `O(mα(n))`    |
// - m is the number of UNION+FIND operations
// - n is the number of elements
// - α(n) is the inverse Ackerman function
type UnionFind struct {
	parents []int
	ranks   []int
}

func New(size int) UnionFind {
	uf := UnionFind{}
	parents := make([]int, size)
	for i := 0; i < size; i++ {
		parents[i] = i
	}
	uf.parents = parents
	uf.ranks = make([]int, size)
	return uf
}

func (uf *UnionFind) Find(x int) int {
	cur := x
	for uf.parents[cur] != cur {
		cur = uf.parents[cur]
	}
	xParent := cur

	// Path Compression
	for uf.parents[x] != xParent {
		next := uf.parents[x]
		uf.parents[x] = xParent
		x = next
	}

	return xParent
}

func (uf *UnionFind) Union(x int, y int) bool {
	xParent, yParent := uf.Find(x), uf.Find(y)

	if xParent == yParent {
		return false
	}

	// Union by Rank
	xRank, yRank := uf.ranks[xParent], uf.ranks[yParent]

	if xRank > yRank {
		// Choose the higher rank tree
		uf.parents[yParent] = xParent
	} else if xRank < yRank {
		// Choose the higher rank tree
		uf.parents[xParent] = yParent
	} else {
		// Choose a tree arbitrarily
		// and increase its rank
		uf.parents[xParent] = yParent
		uf.ranks[yParent]++
	}

	return true
}
