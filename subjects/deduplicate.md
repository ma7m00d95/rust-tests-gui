## Exercise 05: Deduplication

Remove all repeated elements of a list, preserving initial ordering.

```rust
fn deduplicate(list: &mut Vec<i32>);
```

Example:
```rust
let mut v = vec![1, 2, 2, 3, 2, 4, 3];
deduplicate(&mut v);
assert_eq!(v, [1, 2, 3, 4]);
```
