![CI](https://github.com/neutiyoo/data-structures-and-algorithms-in-golang/workflows/CI/badge.svg)

# Data Structures & Algorithms in Golang

## Asymptotic Analysis

**Suppress constant factors and lower-order terms**

<details>
<summary>Why</summary>

<small>

- Constant factors are too system-dependent.
- Lower-order terms are irrelavant for large inputs.

</small>
</details>

<details>
<summary>Examples</summary>

<small>

- `Θ(2n+2)` → `Θ(n)`
- `Θ(6n log n + 6n)` → `Θ(n log n)`
- `Θ(2ⁿ+n⁹)` → `Θ(2ⁿ)`
- `Θ(a+b), where b < a` → `Θ(a)`
- `Θ(a+b)` → `Θ(a+b)`

</small>
</details>

<br/>

**Asymptotic notation**

| Notation | Formal Definition                                                                                  |
| :------: | :------------------------------------------------------------------------------------------------- |
|  Big-θ   | `θ(g(n)) = f(n) ⟺ ∃ positive constants c₁, c₂, and n₀ such that 0 ≤ c₁g(n) ≤ fn) ≤ c₂g(n) ∀n ≥ n₀` |
|  Big-O   | `O(g(n)) = f(n) ⟺ ∃ positive constants c, n₀ such that 0 ≤ f(n) ≤ cg(n) ∀n ≥ n₀`                   |
|  Big-Ω   | `Ω(g(n)) = f(n) ⟺ ∃ positive constants c, n₀ such that 0 ≤ cg(n) ≤ f(n) ∀n ≥ n₀`                   |

## Data Structures

- [Linked List](internal/linked_list)
- [Stack](internal/stack)
- [Queue](internal/queue)
- [Binary Tree](internal/tree)
- [Binary Search Tree](internal/tree/bst)
- [Heap](internal/heap)
- [Trie](internal/trie)
- [Union-Find](internal/union_find)
- [Graph](internal/graph)
- [Hash Table](internal/hash_table)
- [Bloom Filter](internal/bloom_filter)

## Search Algorithms

- [Binary Search](internal/sort/search)

## Sorting Algorithms

- [Merge Sort](internal/sort/merge_sort)
- [Quicksort](internal/sort/quicksort)
- [Heap Sort](internal/sort/heap_sort)
- [Insertion Sort](internal/sort/insertion_sort)
- [Selection Sort](internal/sort/selection_sort)
- [Bubble Sort](internal/sort/bubble_sort)

## Graph Algorithms

- [Breadth-First Search](internal/graph/bfs)
- [Depth-First Search](internal/graph/dfs)
- [Topological Sort](internal/graph/topological_sort)
- [Dijkstra's Algorithm](internal/graph/dijkstra)
- [Prim's Algorithm](internal/graph/prim)
- [Kruskal's Algorithm](internal/graph/kruskal)

## References

- [Introduction to Algorithms Third Edition](https://mitpress.mit.edu/books/introduction-algorithms-third-edition), Thomas H. Cormen
- [Algorithms Illuminated: Part 1: The Basics](https://www.amazon.com/dp/0999282905), Tim Roughgarden
- [Algorithms Illuminated (Part 2): Graph Algorithms and Data Structures](https://www.amazon.com/dp/0999282921), Tim Roughgarden
- [Algorithms Illuminated (Part 3): Greedy Algorithms and Dynamic Programming](https://www.amazon.com/dp/0999282948), Tim Roughgarden
