## Exercise 04: Command Multiplexer

Run multiple commands in parallel and gather their output.

Example:
```txt
>_ cargo run -- echo a , sleep 1 , ls
```

* Output must be displayed entirely as soon as a child finishes.
* Standard error must be ignored.
