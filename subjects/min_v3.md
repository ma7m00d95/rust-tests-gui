## Exercise 01: Point Of No Return (v3)

Create a generic `min` function for any type that supports `PartialOrd`.

```rust
fn min<T: PartialOrd>(a: T, b: T) -> T;
```

* Still not allowed to use `return`!
