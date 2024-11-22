# go-slugger

`go-slugger` is a simple yet powerful Go package for generating slugs from strings. It handles various languages, special characters, and customizable delimiters, making it perfect for generating clean, SEO-friendly URLs and slugs for any project.

## Project Stats
![Forks](https://img.shields.io/github/forks/jmrashed/go-slugger?style=social)
![Issues](https://img.shields.io/github/issues/jmrashed/go-slugger)
![Releases](https://img.shields.io/github/downloads/jmrashed/go-slugger/latest/total)


## Features

- Convert strings to slugs in multiple languages.
- Customizable delimiters (default is `-`).
- Option to remove diacritics (accents).
- Handles special characters and unwanted symbols.
- Easy to integrate with any Go project.

## Installation

To install `go-slugger`, use Go modules to include it in your project:

```bash
go get github.com/jmrashed/go-slugger
```

## Usage

Once installed, you can use `go-slugger` to generate slugs easily.

### Basic Example

```go
package main

import (
	"fmt"
	"github.com/jmrashed/go-slugger"
)

func main() {
	// Create a new Slugger instance with default settings
	slugger := slugger.NewSlugger("-", true)

	// Generate a slug
	slug := slugger.Slugify("Hello World! This is a test.")
	fmt.Println(slug) // Output: hello-world-this-is-a-test
}
```

### Custom Delimiters and Diacritic Removal

You can customize the delimiter and choose whether to remove diacritical marks (accents).

```go
package main

import (
	"fmt"
	"github.com/jmrashed/go-slugger"
)

func main() {
	// Create a Slugger instance with custom delimiter and remove diacritics
	slugger := slugger.NewSlugger("_", false)

	// Generate a slug
	slug := slugger.Slugify("Café au lait")
	fmt.Println(slug) // Output: café_au_lait
}
```

### Example of Custom Delimiters with Special Characters

```go
package main

import (
	"fmt"
	"github.com/jmrashed/go-slugger"
)

func main() {
	// Use custom delimiter and remove unwanted characters
	slugger := slugger.NewSlugger("_", true)

	// Generate a slug
	slug := slugger.Slugify("My--Custom--Slugger!2024")
	fmt.Println(slug) // Output: my_custom_slugger_2024
}
```

## API

### `func NewSlugger(delimiter string, removeDiacritics bool) *Slugger`

Creates a new `Slugger` instance with the specified delimiter and diacritic removal option.

- `delimiter` (string): The delimiter used to replace spaces and non-alphanumeric characters (default is `"-"`).
- `removeDiacritics` (bool): Whether to remove diacritical marks (accents) from characters (default is `true`).

### `func (s *Slugger) Slugify(input string) string`

Converts the input string to a slug, based on the options set when creating the `Slugger`.

- **Returns**: The generated slug as a string.

## Testing

You can run the tests for the package using:

```bash
go test
```

We use the [Testify](https://github.com/stretchr/testify) testing framework for assertion-based testing. Ensure you have all dependencies installed by running:

```bash
go get github.com/stretchr/testify
```

### Example Tests

```go
package slugger

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSlugger(t *testing.T) {
	// Test case 1: Basic Slugify with default settings
	slugger := NewSlugger("-", true)
	result := slugger.Slugify("Hello World! This is a test.")
	expected := "hello-world-this-is-a-test"
	assert.Equal(t, expected, result)

	// Test case 2: Slugify without diacritic removal
	slugger = NewSlugger("-", false)
	result = slugger.Slugify("Café au lait")
	expected = "café-au-lait"
	assert.Equal(t, expected, result)

	// Test case 3: Slugify with custom delimiter
	slugger = NewSlugger("_", true)
	result = slugger.Slugify("Hello World! This is a test.")
	expected = "hello_world_this_is_a_test"
	assert.Equal(t, expected, result)
}
```

## Contributing

We welcome contributions! If you have a feature request or found a bug, feel free to open an issue or submit a pull request.

### Steps for contributing:

1. Fork this repository.
2. Clone your fork to your local machine.
3. Create a new branch for your feature or bugfix.
4. Make your changes and test them.
5. Submit a pull request with a clear description of the changes you made.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 