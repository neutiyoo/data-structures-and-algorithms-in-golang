# Hash Tables

## Operations

| Name   | Time Complexity (average) | Time Complexity (worst) |
| ------ | :-----------------------: | :---------------------: |
| INSERT |          `O(1)`           |         `O(n)`          |
| LOOKUP |          `O(1)`           |         `O(n)`          |
| DELETE |          `O(1)`           |         `O(n)`          |

## Resolving Collisions

- Separate Chaining
- Open Addressing
  - Linear Probing
  - Double Hashing

## Good Hash Table Performance

**The load factor of a hash table**
`α := # of objects in hash table / # of buckets of hash table`

1. Keep the load factor under control

   - `α = O(1)` is necessary condition for operations to run in constant time
   - With open addressing, `α << 1`

2. Use good hash function

   - guaranteed to spread every data set out evenly
   - easy to store / very fast to evaluate
