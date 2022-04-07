# Bloom Filters

| Required Space |
| :------------: |
|     `O(m)`     |

- `m` is the number of bits

## Operations

| Name   | Time Complexity |
| ------ | :-------------: |
| INSERT |     `O(k)`      |
| LOOKUP |     `O(k)`      |

- `k` is the number of hash function

## Comparison to Hash Tables

Pros

- More space efficient

Cons

- Can’t store an associated object
- No deletions
- Small false positive probability
