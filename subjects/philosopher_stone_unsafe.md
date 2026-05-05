## Exercise 01: Philospher's Stone

Use `transmute` to convert between different types while preserving bit patterns.

```rust
impl PhilosopherStone {
    fn transmute_iron(self, iron: Iron) -> [GoldNugget; 2];
    fn transmute_metal<M: Metal>(self, metal: &M) -> &Gold;
}
```
