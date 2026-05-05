## Exercise 06: Foreign User

Link Rust code with an external C library using FFI.

```rust
extern "C" {
    fn create_database(database: *mut t_database) -> e_result;
}

impl Database {
    fn create_user(&mut self, name: &CStr) -> Result<Id, Error>;
}
```
