## Exercise 02: Where's My Pizza?

Implement a `PizzaStatus` enum with methods to predict delivery time.

```rust
enum PizzaStatus {
    Ordered,
    Cooking,
    Cooked,
    Delivering,
    Delivered,
}

impl PizzaStatus {
    fn from_delivery_time(ordered_days_ago: u32) -> Self;
    fn get_delivery_time_in_days(&self) -> u32;
}
```
