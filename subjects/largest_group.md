## Exercise 03: Largest Group

Write a **function** that returns the largest subslice of `haystack` that contains *all* numbers in `needle`.

```rust
fn largest_group(haystack: &[u32], needle: &[u32]) -> &[u32];
```

* When multiple groups match, the largest one is returned.
* When multiple largest groups are found, the first one is returned.
