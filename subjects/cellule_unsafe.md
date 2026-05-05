## Exercise 03: Cellule<T>

Recreate `Cell<T>` as `Cellule<T>` using `UnsafeCell`.

```rust
impl<T> Cellule<T> {
    fn set(&self, value: T);
    fn replace(&self, value: T) -> T;
    fn get(&self) -> T where T: Copy;
}
```
