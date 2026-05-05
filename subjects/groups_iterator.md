## Exercise 05: From Side To Side

Create a `Groups` iterator that splits a string based on a predicate function `f`.

```rust
impl<'a, F> Groups<'a, F> {
    pub fn new(s: &'a str, f: F) -> Self
    where
        F: FnMut(char) -> bool;
}
```

Must implement `DoubleEndedIterator` for reverse support.
