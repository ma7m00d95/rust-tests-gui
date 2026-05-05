# Rust Piscine: Local Practice Guide

This guide explains how to use the `rust-tests` repository to practice for the Reboot01/01Edu Rust Piscine without using Docker.

---

## 1. Local Environment Setup
Run these commands once in your terminal to ensure your local Rust installation matches the platform's requirements:

```bash
# Install the formatter and linter, be careful not everything is working, you need to check how to install it

rustup component add rustfmt
rustup component add clippy
```

---

## 2. Folder Overview

| Folder | Purpose |
| :--- | :--- |
| `solutions/` | **Your Reference:** Contains the completed code for every exercise. Use this if you get stuck. |
| `tests/` | **The Grader:** Contains the test files that check your code. |
| `tests_utility/` | **Internal:** Required for tests to function. Do not modify. |

---

## 3. How to Practice
To simulate the piscine experience locally, follow these steps:

1.  **Select an Exercise**: Pick a folder from `solutions/` (e.g., `hello`).
2.  **Write your Code**: Modify the code inside `solutions/[exercise_name]/src/main.rs` (or `lib.rs`) with your own solution.
3.  **Run the Test**: From the **root** of the `rust-tests` directory, run the grading script:
    ```bash
    bash tests/test_exercises.sh [exercise_name]
    ```
    *Example:* `bash tests/test_exercises.sh hello`

---

## 4. Useful Commands

| Command | Result |
| :--- | :--- |
| `bash tests/test_exercises.sh -v [name]` | Run test with **Verbose** (detailed) feedback. |
| `bash tests/test_exercises.sh -h` | Show all available script options/help. |
| `cargo test` | Run manually while inside a `tests/[name]_test` folder. |

---

> **Note:** The order of the projects in the `solutions` folder is alphabetical. For the actual 4-week timeline, follow the **Quest 01 → Quest 09** progression.
