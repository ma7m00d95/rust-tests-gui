## Exercise 02: Last Error

Use `thread_local!` to store the last error of a thread.

```rust
enum Error { Success, FileNotFound, ... }

impl Error {
    fn last() -> Self;
    fn make_last(self);
}
```
