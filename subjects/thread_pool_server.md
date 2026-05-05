## Exercise 06: 404 Not Found

Implement a `ThreadPool` and use it to build a multi-threaded HTTP server.

```rust
struct ThreadPool {
    threads: Vec<JoinHandle<()>>,
    task_sender: Sender<Task>,
}
```

The server should respond with "404 Not Found" to every connection.
