## Exercise 04: Boxes Into Boxes

Sort a list of boxes (`[width, height]`) so each box is contained in the previous one. Panic if impossible.

```rust
fn sort_boxes(boxes: &mut [[u32; 2]]);
```

Example:
```rust
let mut boxes = [[3, 3], [4, 3], [1, 0], [5, 7], [3, 3]];
sort_boxes(&mut boxes);
assert_eq!(boxes, [[5, 7], [4, 3], [3, 3], [3, 3], [1, 0]]);
```
