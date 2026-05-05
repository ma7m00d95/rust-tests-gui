## Exercise 02: Carton

Create a `Carton<T>` smart pointer that manages a single heap allocation.

```rust
struct Carton<T> { /* ... */ }

impl<T> Carton<T> {
    fn new(value: T) -> Self;
    fn into_inner(this: Self) -> T;
}
```

* Ensure correct variance and drop tracking.
* Manually manage `alloc` and `dealloc`.
