# Contributing

First, thank you for contributing! We love and encourage pull requests from everyone.

### Abstract

Before submitting major changes, here are a few guidelines to follow:

#### 1. Check for existing issues or PRs

Before starting, please check the [open issues][issues] and [pull requests][prs] for existing discussions. If you don't find an issue or PR, please open a new issue to discuss the new feature or enhancement.

#### 2. Open an issue

If you plan to add a new feature or enhancement, please open an issue first to discuss it. This will help us to understand your needs and provide feedback.

#### 3. Branching

Always add features or fixes in a new branch, and never in `main`. This will help us to review and test your changes before merging them into the main codebase.

#### 4. Write tests

We have a test suite for the project, and we expect you to write tests for your changes. Make sure the test suite passes locally and on CI.

#### 5. Run `make pre`

Before submitting your changes, run `make pre` to ensure your changes pass all checks. This will run the linter, tests, and other checks.

#### 6. Update the documentation

Make sure to update the documentation to reflect your changes. This includes the README, and any other relevant documentation.

#### 7. Open a pull request

[issues]: https://github.com/thuongtruong1009/gouse/issues
[prs]: https://github.com/thuongtruong1009/gouse/pulls

### Commands

### 1. Install dependencies

```bash
$ make install
```

#### 3. Run

```bash
$ make run
```

#### 4. Build

```bash
$ make build
```

#### 5. Run tests

```bash
$ make test
```

#### 6. Run linters & formatters

```bash
$ make lint
```

#### 7. Run benchmarks

```bash
$ make bench
```

#### 8. Auto-generate docs

```bash
$ make doc
```

#### 9. Clean up the project

```bash
$ make clean
```

#### 10. Pre-commit checks

```bash
$ make pre
```

#### 11. Get help

```bash
$ make help
```

### License

By contributing, you agree that your contributions will be licensed under its [MIT License](../LICENSE).

### References

- [GitHub Flow](https://guides.github.com/introduction/flow/)
- [How to Contribute to Open Source](https://opensource.guide/how-to-contribute/)
- [How to Write the Perfect Pull Request](https://github.blog/2015-01-21-how-to-write-the-perfect-pull-request/)
