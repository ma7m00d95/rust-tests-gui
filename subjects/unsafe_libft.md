## Exercise 00: Libft

Recreate basic string functions using unsafe pointers.

```rust
fn ft_swap<T>(a: &mut T, b: &mut T);
unsafe fn ft_strlen(s: *const u8) -> usize;
unsafe fn ft_strcpy(dst: *mut u8, src: *const u8);
```

**Rule**: Every `unsafe` block must have a `# Safety` documentation section.
