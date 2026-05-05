## Exercise 04: Atomical

Create a `Unique` type using `AtomicU8` to ensure every instance has a unique ID.

```rust
struct Unique(u8);

impl Unique {
    pub fn new() -> Self;
}
```
