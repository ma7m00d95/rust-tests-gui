## Exercise 02: The Name Of Colors

```rust
const fn color_name(color: &[u8; 3]) -> &str;
```

Rules:
* `[0, 0, 0]` is "pure black".
* `[255, 255, 255]` is "pure white".
* `[255, 0, 0]` is "pure red".
* `[0, 255, 0]` is "pure green".
* `[0, 0, 255]` is "pure blue".
* `[128, 128, 128]` is "perfect grey".
* Components < 31 is "almost black".
* Red > 128, others 0-127 is "redish".
* Green > 128, others 0-127 is "greenish".
* Blue > 128, others 0-127 is "blueish".
* Otherwise is "unknown".
