## Exercuse 04: Monotically Increasing

Create a type named `Increasing<I>` that filters any non-strictly-increasing items from an iterator.

```rust
struct Increasing<I: Iterator> {
    inner: I,
    last: Option<I::Item>,
}
```
