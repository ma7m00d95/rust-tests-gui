## Exercise 06: Lexial Analysis

Create a simple token parser using an `enum Token`.

```rust
enum Token<'a> {
    Word(&'a str),
    RedirectStdout,
    RedirectStdin,
    Pipe,
}

fn next_token(s: &mut &str) -> Option<Token>;
```
