## Exercise 07: String Pattern Compare

```rust
fn strpcmp(query: &[u8], pattern: &[u8]) -> bool;
```

* `strpcmp` determines whether `query` matches the given `pattern`.
* `pattern` may optionally contain `"*"` characters, which can match any number of any character in
the query string.

Example:

```rust
assert!(strpcmp(b"abc", b"abc"));

assert!(strpcmp(b"abcd", b"ab*"));
assert!(!strpcmp(b"cab", b"ab*"));

assert!(strpcmp(b"dcab", b"*ab"));
assert!(!strpcmp(b"abc", b"*ab"));

assert!(strpcmp(b"ab000cd", b"ab*cd"));
assert!(strpcmp(b"abcd", b"ab*cd"));

assert!(strpcmp(b"", b"****"));
```
