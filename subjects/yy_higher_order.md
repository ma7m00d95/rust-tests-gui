## Exercise 01: YYYYYYYYYYYYYY

Implement three functions using closures for sequence processing.

```rust
fn collayz<F: FnMut(u32)>(start: u32, f: F);
fn yes<F: FnOnce() -> String>(f: F) -> !;
fn print_byes<F: FnMut() -> Option<u8>>(f: F);
```
