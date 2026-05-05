## Exercise 07: Rendez-Vous

Create a "Rendez-Vous" primitive where two threads exchange values.

```rust
struct RendezVous<T> { /* ... */ }

impl<T> RendezVous<T> {
    fn wait(&self, value: T) -> T;
}
```

`wait` blocks until a second thread calls it, then they swap their values.
