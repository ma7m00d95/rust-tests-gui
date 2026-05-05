## Exercise 06: More Format

Implement a custom formatting system using the `Print` trait.

```rust
trait Print {
    fn print(&self, write: &mut WriteFn) -> Result<(), FormatError>;
}

fn format_string(s: &str, values: &[&dyn Print]) -> Result<String, FormatError>;
```
