## Exercise 06: A Singly-Linked List

Implement a generic singly-linked list `List<T>`.

```rust
struct List<T> {
    head: Option<Box<Node<T>>>
}

impl<T> List<T> {
    fn push_front(&mut self, value: T);
    fn push_back(&mut self, value: T);
    fn count(&self) -> usize;
    fn get(&self, i: usize) -> Option<&T>;
}
```
