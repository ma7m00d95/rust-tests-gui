## Exercise 05: All Over Me

Implement color blending logic.

```rust
struct Color {
    red: u8,
    green: u8,
    blue: u8,
}

impl Color {
    fn closest_mix(self, palette: &[(Self, u8)], max: u32) -> Self;
}
```

Formula: `C = A * alpha + B * (1 - alpha)`
