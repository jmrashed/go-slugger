
         _              _               _            _       _                  _              _              _            _      
        /\ \           /\ \            / /\         _\ \    /\_\               /\ \           /\ \           /\ \         /\ \    
       /  \ \         /  \ \          / /  \       /\__ \  / / /         _    /  \ \         /  \ \         /  \ \       /  \ \   
      / /\ \_\       / /\ \ \        / / /\ \__   / /_ \_\ \ \ \__      /\_\ / /\ \_\       / /\ \_\       / /\ \ \     / /\ \ \  
     / / /\/_/      / / /\ \ \      / / /\ \___\ / / /\/_/  \ \___\    / / // / /\/_/      / / /\/_/      / / /\ \_\   / / /\ \_\ 
    / / / ______   / / /  \ \_\     \ \ \ \/___// / /        \__  /   / / // / / ______   / / / ______   / /_/_ \/_/  / / /_/ / / 
   / / / /\_____\ / / /   / / /      \ \ \     / / /         / / /   / / // / / /\_____\ / / / /\_____\ / /____/\    / / /__\/ /  
  / / /  \/____ // / /   / / /   _    \ \ \   / / / ____    / / /   / / // / /  \/____ // / /  \/____ // /\____\/   / / /_____/   
 / / /_____/ / // / /___/ / /   /_/\__/ / /  / /_/_/ ___/\ / / /___/ / // / /_____/ / // / /_____/ / // / /______  / / /\ \ \     
/ / /______\/ // / /____\/ /    \ \/___/ /  /_______/\__\// / /____\/ // / /______\/ // / /______\/ // / /_______\/ / /  \ \ \    
\/___________/ \/_________/      \_____\/   \_______\/    \/_________/ \/___________/ \/___________/ \/__________/\/_/    \_\/ 

GoSlugger version v1.0.1 2024-11-22

# go-slugger

`go-slugger` is a Go package for generating slugs from strings. It supports multiple languages, custom delimiters, diacritic removal, and more. This package is useful for generating clean, SEO-friendly slugs for URLs, making it easy to integrate with your web applications, CMS, or any system that needs slug generation.

[![Forks](https://img.shields.io/github/forks/jmrashed/go-slugger?style=social)](https://github.com/jmrashed/go-slugger/forks)
[![Issues](https://img.shields.io/github/issues/jmrashed/go-slugger)](https://github.com/jmrashed/go-slugger/issues)
[![Releases](https://img.shields.io/github/downloads/jmrashed/go-slugger/latest/total)](https://github.com/jmrashed/go-slugger/releases)
![License](https://img.shields.io/github/license/jmrashed/go-slugger)

## Features

- **Multi-language Support**: Generates slugs for strings in multiple languages, converting accented characters (e.g., `é`, `ö`) to their nearest ASCII equivalents.
- **Custom Delimiters**: Allows you to specify a custom delimiter (default is `-`) for separating words in the slug.
- **Diacritic Removal**: Automatically removes diacritical marks (accents) from characters.
- **Truncation**: Optionally truncate slugs to a specified length.
- **Case Sensitivity**: Control over whether the generated slug is lowercase.

## Installation

To install `go-slugger`, use the Go Modules feature to fetch the package:

```bash
go get github.com/jmrashed/go-slugger
```

## Usage

### Basic Usage

To start using `go-slugger`, import the package into your Go file and use the `GenerateSlug` function.

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

### Custom Delimiters

You can specify a custom delimiter (default is `-`) to separate words in the slug.

```go
slug := slugger.GenerateSlug("Hello World!", "_")
fmt.Println(slug) // Output: hello_world
```

### Diacritic Removal

The package automatically removes diacritical marks (accents) for characters like `é`, `ö`, etc.

```go
slug := slugger.GenerateSlug("Café de Paris")
fmt.Println(slug) // Output: cafe-de-paris
```

### Case Sensitivity

The default behavior is to generate slugs in lowercase. If you want to preserve the original case, you can pass a `false` flag for the `lowercase` parameter.

```go
slug := slugger.GenerateSlug("Hello World!", "-", false)
fmt.Println(slug) // Output: Hello-World
```

### Truncating Slugs

You can specify the maximum length of the generated slug. If the length exceeds this, it will be truncated.

```go
slug := slugger.GenerateSlug("This is a very long title", "-", true, 20)
fmt.Println(slug) // Output: this-is-a-very-lo
```

## API Reference

### `GenerateSlug(input string, delimiter string, lowercase bool, maxLength int) string`

- `input` (string): The string to convert into a slug.
- `delimiter` (string, optional): The delimiter used in the slug. Default is `-`.
- `lowercase` (bool, optional): Whether to generate the slug in lowercase (default: `true`).
- `maxLength` (int, optional): Maximum length of the slug. The slug will be truncated if it exceeds this length.

Returns the generated slug as a string.

## Example

### Full Example

```go
package main

import (
    "fmt"
    "github.com/jmrashed/go-slugger"
)

func main() {
    input := "München ist schön!"
    
    // Generate slug with default settings
    slug1 := slugger.GenerateSlug(input)
    fmt.Println(slug1) // Output: muenchen-ist-schoen
    
    // Generate slug with custom delimiter and case preservation
    slug2 := slugger.GenerateSlug(input, "_", false)
    fmt.Println(slug2) // Output: Muenchen_ist_schoen
    
    // Generate slug with truncation and custom length
    slug3 := slugger.GenerateSlug(input, "-", true, 15)
    fmt.Println(slug3) // Output: muenchen-ist-s
}
```

## Contribution

We welcome contributions to `go-slugger`! If you’d like to contribute, please follow the steps below:

1. Fork the repository
2. Clone your fork locally
3. Create a new branch for your changes
4. Make your changes and add tests
5. Open a pull request to the `main` branch

For more details, refer to the [CONTRIBUTING.md](CONTRIBUTING.md) file.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Changelog

For a detailed list of changes, refer to the [CHANGELOG.md](CHANGELOG.md).

## Support

If you have any questions or need support, feel free to open an issue in the [GitHub Issues](https://github.com/jmrashed/go-slugger/issues) section.

For more detailed documentation, visit the [documentation folder](docs/). 