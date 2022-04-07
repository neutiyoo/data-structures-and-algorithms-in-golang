# Union-Find

| Operation                                    | Time Complexity |
| -------------------------------------------- | :-------------: |
| UNION+FIND (union by rank, path compression) |   `O(mα(n))`    |

- `m` is the number of UNION+FIND operations
- `n` is the number of elements
- `α(n)` is the inverse Ackerman function

## Heuristics to Improve the Running Time

### Union by Rank

| ![union by rank](/asset/img/union_find/union_find.001.png) |
| :--------------------------------------------------------: |
|                     _X Rank < Y Rank_                      |

| ![union by rank](/asset/img/union_find/union_find.002.png) |
| :--------------------------------------------------------: |
|                     _X Rank == Y Rank_                     |

### Path Compression

| ![parh compression](/asset/img/union_find/union_find.003.png) |
| :-----------------------------------------------------------: |
|                      _Path Compression_                       |
