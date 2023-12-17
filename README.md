<div align="center">
    <img src="https://github.com/thuongtruong1009/gouse/actions/workflows/analysis.yml/badge.svg?branch=main" alt="ci_status">
    <img src="https://github.com/thuongtruong1009/gouse/actions/workflows/ci.yml/badge.svg?branch=main" alt="ci_status">
    <a href="https://sonarcloud.io/summary/new_code?id=thuongtruong1009_gouse"><img src="https://sonarcloud.io/api/project_badges/measure?project=thuongtruong1009_gouse&metric=alert_status" alt="sonar"></a>
    <a href="https://app.codacy.com/gh/thuongtruong1009/gouse/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade"><img src="https://app.codacy.com/project/badge/Grade/21f940894abd4e0384ef8b84adc294da" alt="codacy"></a>
    <a href="https://app.deepsource.com/gh/thuongtruong1009/gouse/" target="_blank"><img alt="DeepSource" title="DeepSource" src="https://app.deepsource.com/gh/thuongtruong1009/gouse.svg/?label=resolved+issues&show_trend=false&token=VqBk6AsowWePd3khy3AUkvXJ"/></a>
    <a href="https://goreportcard.com/report/thuongtruong1009/gouse"><img src="https://goreportcard.com/badge/github.com/thuongtruong1009/gouse" alt="go_report_card"></a>
    <a href="https://codecov.io/gh/thuongtruong1009/gouse"><img src="https://codecov.io/gh/thuongtruong1009/gouse/branch/main/graph/badge.svg" alt="codecov"></a>
    <a href="https://dl.circleci.com/status-badge/redirect/gh/thuongtruong1009/gouse/tree/main"><img src="https://dl.circleci.com/status-badge/img/gh/thuongtruong1009/gouse/tree/main.svg?style=svg" alt="circleci"></a>
    <a href="https://pkg.go.dev/github.com/thuongtruong1009/gouse"><img src="https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square" alt="go.dev"></a>
    <img src="./public/count.svg" alt="gouse_functions_count">
    <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/thuongtruong1009/gouse">
    <a href=""><img src="https://img.shields.io/github/license/thuongtruong1009/gouse" alt="license"></a>
    <!-- <a href="https://app.codacy.com/gh/thuongtruong1009/gouse/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_Coverage"><img src="https://app.codacy.com/project/badge/Coverage/21f940894abd4e0384ef8b84adc294da" alt="codacy"></a> -->
    <!-- <a href="https://sourcegraph.com/github.com/thuongtruong1009/gouse?badge"><img src="https://sourcegraph.com/github.com/thuongtruong1009/gouse/-/badge.svg" alt="sourcegraph"></a> --> 
</div>

# ![](/public/banner.png)

## 📝 Why Gouse?

- Gouse is a modern essential [`Golang`](https://golang.org/) utility package delivering consistency, modularity, performance, & extras presets inspired by [`Javascript`](https://developer.mozilla.org/en-US/docs/Web/JavaScript) and [`Lodash`](https://lodash.com/).
- Gouse provides a wide variety of methods, taking the hassle out of working with arrays, numbers, objects, strings, etc.
- Each method has a different set of features, so you can pick the ones that fit your project best.
- Easy to use, chainable, and extendable, and available in a variety of builds & module formats.
- `Javascript` and `Redis` user-friendly syntax.
- No config - import as utility functions.
- Available functions to handle complex logic usecases such as database connection, build APIs, error handling, log management... can all be set up quickly and easily.

## ✨ Motivation

- Go has emerged as a server language, but it still doesn't have complete and consistent packages available to support coding development.

- To address that need, Gouse was created as a powerful toolkit for Go developers, a collection of built-in functions and best practices that provide comprehensive, powerful, and reliable solutions. Trusted to build services, APIs, and web applications.

## 🚀 Features

> Below is a list of modules that Gouse supports. This project is still in development, so not all features are available.

- [ ] API
- [x] [`Array`](docs/gen/sample/array.md)
- [x] [`Console`](docs/gen/sample/console.md)
- [ ] Collection
- [ ] Connection
- [x] [`Date`](docs/gen/sample/date.md)
- [ ] Error
- [x] [`Function`](docs/gen/sample/function.md)
- [x] [`Helper`](docs/gen/sample/helper.md)
- [x] [`I/O`](docs/gen/sample/io.md)
- [ ] Lang
- [ ] Log
- [x] [`Math`](docs/gen/sample/math.md)
- [x] [`Net`](docs/gen/sample/net.md)
- [x] [`Number`](docs/gen/sample/number.md)
- [x] [`Regex`](docs/gen/sample/regex.md)
- [x] [`Struct`](docs/gen/sample/struct.md)
- [x] [`String`](docs/gen/sample/string.md)
- [ ] Tools
- [x] [`Type`](docs/gen/sample/type.md)

## 📋 Requirements

> Compatibility with Go >= `1.18`

## 📦 Installation

```go
go get github.com/thuongtruong1009/gouse
```

## 🕯️ Quick Start

```go
package main

import "github.com/thuongtruong1009/gouse"

func main() {
    gouse.Stater()
}
```

## 🦄 Usage

- Using package directly in your module as ultra-lightweight utility functions.

```go
package main

import (
    "fmt"
    "github.com/thuongtruong1009/gouse/math"
)

func main() {
    fmt.Println(math.Add(1, 2))
}
```

- View more examples at [`sample`](sample) folder.

## 📖 Documentation

- Package documentation reference at [![Dev package](https://pkg.go.dev/badge/github.com/thuongtruong1009/gouse)](https://pkg.go.dev/github.com/thuongtruong1009/gouse)

<!-- ## 📊 Benchmark

## 📚 Examples -->

## 🛠️ Development

#### 1. Clone this repository

```bash
$ git clone https://github.com/thuongtruong1009/gouse.git
```

#### 2. Install dependencies

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

## 📝 Contributing

- We welcome your contributions! If you're looking for issues to work on, try looking at the good first issue list. We do our best to tag issues suitable for new external contributors with that label, so it's a great way to find something you can help with!

- [`Code of Conduct`](.github/CODE_OF_CONDUCT.md) and [`Contributing Guide`](.github/CONTRIBUTING.md) for the project.

## 📄 License

- Gouse is released under the [`MIT License`](LICENSE). See the LICENSE file for more information.

- For more information, see the [`Licensing FAQs`](https://opensource.org/faq#mit-vs-bsd).

## 📌 Support

- The tool has been tested on a variety of inputs, but it's not perfect.
- For support in using Gouse, please reach out in the following venues:
  - [`Raise Issues`](https://github.com/thuongtruong1009/gouse/issues/new) - For generally applicable issues and feedback.
  - [`Join Discussions`](https://github.com/thuongtruong1009/gouse/discussions) - For ideas, questions, or issues regarding Gouse's design, development, and future.

## 📜 Changelog

- [`CHANGELOG`](CHANGELOG.md) for project history.

## 🌸 Sponsor

- If you like this project, you can sponsor me on:

[![Github sponsor](https://img.shields.io/badge/GitHub_Sponsors-000000?style=for-the-badge&logo=github&logoColor=white)](https://sponsor.com/thuongtruong1009)
[![Paypal](https://img.shields.io/badge/PayPal-00457C?style=for-the-badge&logo=paypal&logoColor=white)](https://paypal.me/thuongtruong1009)
[![Kofi](https://img.shields.io/badge/ko--fi-F16061?style=for-the-badge&logo=ko-fi&logoColor=white)](https://ko-fi.com/thuongtruong1009)

<!-- ![Static Badge](https://img.shields.io/badge/Buy_Me_A_Coffee-FFDD00?style=for-the-badge&logo=buy-me-a-coffee&logoColor=black) -->

## 📮 Contact

- If you have any questions, please contact me:

[![Email](https://img.shields.io/badge/Gmail-D14836?style=for-the-badge&logo=gmail&logoColor=white)](mailto:thuongtruongofficial@gmail.com)
[![Github](https://img.shields.io/badge/Github-000000?style=for-the-badge&logo=github&logoColor=white)](https://github.com/thuongtruong1009)

<!-- dependencies used -->

## 🧬 Dependencies

Special thanks to the following dependencies that helped make this project possible:

- [github.com/google/uuid](github.com/google/uuid) - A fast and simple UUID library for Go 🔑
- [github.com/AlecAivazis/survey/v2](github.com/AlecAivazis/survey/v2) - A golang library for building interactive prompts with full support for windows and posix terminals 🙋
- [github.com/charmbracelet/bubbletea](github.com/charmbracelet/bubbletea) - A powerful little TUI framework 🏗

## 📚 References

- [Golang](https://golang.org/)
- [Golang Docs](https://pkg.go.dev/)
- [Lodash.js](https://lodash.com/)
- [Lodash collection](https://www.geeksforgeeks.org/lodash/?ref=header_search)
- [Javascript methods reference](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference)
- [File handling in Golang](https://www.cloudhadoop.com/2018/11/learn-golang-tutorials-file-directory.html)
- [Golang x Github Actions](https://github.com/brpaz/github-actions-demo-go)
- [Dynamic HTML](https://css-tricks.com/dynamic-page-replacing-content)
