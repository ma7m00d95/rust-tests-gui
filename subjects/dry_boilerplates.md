## Exercise 03: Dry Boilerplates

Create a type `MyType` and use the `derive` attribute to make it work with `main`.

```rust
fn main() {
    let instance = MyType::default();
    let other_instance = instance.clone();
    assert_eq!(instance, other_instance);
    assert!(instance >= other_instance && instance <= other_instance);
}
```
