## Exercise 01: Logger

Create a thread-safe `Logger` that buffers messages to reduce system calls.

```rust
struct Logger<W> {
    buffer: Box<[u8]>,
    writer: W,
}

impl<W: io::Write> Logger<W> {
    pub fn log(&mut self, message: &str) -> io::Result<()>;
}
```

Spawn 10 threads and have them all log simultaneously without mixing messages.
