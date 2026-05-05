## Exercise 04: Todo List

Create a simple TODO-List application using `struct TodoList` and `enum Command`.

```rust
struct TodoList {
    todos: Vec<String>,
    dones: Vec<String>,
}

impl TodoList {
    fn display(&self);
    fn add(&mut self, todo: String);
    fn done(&mut self, index: usize);
    fn purge(&mut self);
}
```
