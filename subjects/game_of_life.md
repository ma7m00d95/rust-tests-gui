## Exercise 07: The Game Of Life

Create a program that plays Conway's Game Of Life using a `Board` struct and `Cell` enum.

```rust
struct Board {
    width: usize,
    height: usize,
    cells: Vec<Cell>,
}

impl Board {
    fn new(width: usize, height: usize, percentage: u32) -> Self;
    fn step(&mut self);
    fn print(&self, clear: bool);
}
```
