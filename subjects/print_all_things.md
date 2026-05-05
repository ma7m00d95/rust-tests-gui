## Exercise 00: Print All Things

Create a generic function that prints all items in an iterator using exactly one `for` loop.

```rust
fn print_all_things<I>(i: I)
where
    I: IntoIterator,
    I::Item: std::fmt::Debug;
```

Example:
```rust
print_all_things(0..=5); // [ 0 1 2 3 4 5 ]
```
