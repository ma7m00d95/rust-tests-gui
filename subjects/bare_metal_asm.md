## Exercise 07: Bare Metal

Write a program using `#![no_std]` and `#![no_main]` that uses inline assembly.

```rust
fn ft_putchar(c: u8);
fn ft_exit(code: u8) -> !;
```

* Must use `core::arch::asm`.
* Output `42` and exit with code `42`.
