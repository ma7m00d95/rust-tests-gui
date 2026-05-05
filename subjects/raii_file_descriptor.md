## Exercise 04: RAII

Create a `File` wrapper that automatically closes file descriptors using RAII.

```rust
struct File(Fd);

impl File {
    fn open(file: &CStr) -> Result<Self, Errno>;
    fn write(&self, data: &[u8]) -> Result<usize, Errno>;
}
```

Uses `libc` for system calls and `errno` for error handling.
