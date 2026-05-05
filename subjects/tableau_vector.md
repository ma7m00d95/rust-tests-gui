## Exercise 05: Tableau

Recreate a dynamic vector `Tableau<T>` from scratch using raw pointers.

```rust
impl<T> Tableau<T> {
    const fn new() -> Self;
    fn push(&mut self, item: T);
    fn pop(&mut self) -> Option<T>;
}
```

**Note**: Be extremely careful with memory leaks during panics!
