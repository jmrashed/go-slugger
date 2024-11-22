# go-slugger - Usage Guide

`go-slugger` is a Go package for generating slugs from strings with support for multiple languages, diacritic removal, and customizable delimiters.

This guide provides examples and detailed instructions on how to use the package in your Go projects.

## Installation

To use `go-slugger` in your Go project, simply install it via Go Modules:

```bash
go get github.com/jmrashed/go-slugger
```

## Basic Usage

To start using `go-slugger` to generate slugs, you need to import it in your Go file.

### Example 1: Basic Slug Generation

```go
package main

import (
    "fmt"
    "github.com/jmrashed/go-slugger"
)

func main() {
    input := "Hello World! Welcome to Go."
    slug := slugger.GenerateSlug(input)
    fmt.Println(slug) // Output: hello-world-welcome-to-go
}
```

### Example 2: Custom Delimiters

You can customize the delimiter used in the generated slug (the default delimiter is `-`).

```go
package main

import (
    "fmt"
    "github.com/jmrashed/go-slugger"
)

func main() {
    input := "Hello World! Welcome to Go."
    slug := slugger.GenerateSlug(input, "_") // Use underscore as delimiter
    fmt.Println(slug) // Output: hello_world_welcome_to_go
}
```

### Example 3: Diacritic Removal (Accents)

By default, the package removes diacritics (accents) from characters like `é`, `ö`, etc., to ensure compatibility with URLs.

```go
package main

import (
    "fmt"
    "github.com/jmrashed/go-slugger"
)

func main() {
    input := "Café de Paris"
    slug := slugger.GenerateSlug(input)
    fmt.Println(slug) // Output: cafe-de-paris
}
```

### Example 4: Handle Multiple Languages

`go-slugger` supports a wide range of characters from different languages, converting them into their closest ASCII equivalents.

```go
package main

import (
    "fmt"
    "github.com/jmrashed/go-slugger"
)

func main() {
    input := "München ist schön!"
    slug := slugger.GenerateSlug(input)
    fmt.Println(slug) // Output: muenchen-ist-schoen
}
```

## Advanced Features

### Case Sensitivity

By default, slugs are generated in lowercase. If you need to preserve the case of the original string, you can pass a `false` flag to the `GenerateSlug` function.

```go
slug := slugger.GenerateSlug(input, "-", false) // Preserve case
```

### Truncate Slug

If you need to limit the length of the generated slug (e.g., for SEO purposes), you can truncate the result.

```go
slug := slugger.GenerateSlug(input, "-", true, 20) // Truncate to 20 characters
fmt.Println(slug) // Output: hello-world-welcom
```

## API Reference

### `GenerateSlug(input string, delimiter string, lowercase bool, maxLength int) string`

- `input` (string): The input string you want to convert into a slug.
- `delimiter` (string, optional): The delimiter to use in the slug. Default is `-`.
- `lowercase` (bool, optional): Whether to generate the slug in lowercase (default: `true`).
- `maxLength` (int, optional): Maximum length of the generated slug. If the slug exceeds this length, it will be truncated.

Returns the generated slug as a string.

---

## Contribution Guidelines

We welcome contributions to `go-slugger`. If you have suggestions, bug reports, or pull requests, please read the [CONTRIBUTING.md](../CONTRIBUTING.md) file for more details. 