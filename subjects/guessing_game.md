## Exercise 06: Guessing Game

```txt
turn-in directory:
    ex06/

files to turn in:
    src/main.rs  Cargo.toml

allowed dependencies;
    ftkit

allowed symbols:
    ftkit::read_number  ftkit::random_number
    i32::cmp  std::cmp::Ordering
```

Create a **program** that plays the guessing game.

```txt
>_ cargo run
Me and my infinite wisdom have found an appropriate secret you shall yearn for.
12
This student might not be as smart as I was told. This answer is obviously too weak.
25
Sometimes I wonder whether I should retire. I would have guessed higher.
19
That is right! The secret was indeed the number 19, which you have brilliantly discovered!
```

You can't use the `<`, `>`, `<=`, `>=` and `==` operators!
