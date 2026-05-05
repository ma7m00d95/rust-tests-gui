## Exercise 07: Comma-Separated Values

Create a generic CSV Encoder & Decoder using traits and macros.

```rust
trait Field: Sized {
    fn encode(&self, target: &mut String) -> Result<(), EncodingError>;
    fn decode(field: &str) -> Result<Self, DecodingError>;
}

fn encode_csv<R: Record>(records: &[R]) -> Result<String, EncodingError>;
fn decode_csv<R: Record>(contents: &str) -> Result<Vec<R>, DecodingError>;
```
