## Exercise 03: 42

Define and implement the `FortyTwo` trait.

```rust
trait FortyTwo {
    fn forty_two() -> Self;
}

fn print_forty_two<T: Debug + FortyTwo>();
```
