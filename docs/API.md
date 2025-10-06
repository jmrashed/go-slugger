# API Documentation

## GenerateSlug Function

### Signature
```go
func GenerateSlug(input string, args ...interface{}) string
```

### Parameters
- `input` (string): The string to convert into a slug
- `args` (variadic): Optional parameters in order:
  1. `delimiter` (string): Separator for words (default: "-")
  2. `lowercase` (bool): Convert to lowercase (default: true)  
  3. `maxLength` (int): Maximum slug length (default: 0 = no limit)

### Examples
```go
// Basic usage
slug := slugger.GenerateSlug("Hello World")
// Output: "hello-world"

// Custom delimiter
slug := slugger.GenerateSlug("Hello World", "_")
// Output: "hello_world"

// Case preservation
slug := slugger.GenerateSlug("Hello World", "-", false)
// Output: "Hello-World"

// Truncation
slug := slugger.GenerateSlug("Very Long Title", "-", true, 10)
// Output: "very-long"
```

## Legacy API (Backward Compatibility)

### Slugger Struct
```go
type Slugger struct {
    Delimiter        string
    RemoveDiacritics bool
}
```

### NewSlugger Function
```go
func NewSlugger(delimiter string, removeDiacritics bool) *Slugger
```

### Slugify Method
```go
func (s *Slugger) Slugify(input string) string
```