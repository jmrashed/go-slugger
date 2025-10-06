# Release Notes v1.1.0

## 🎉 New Features
- **API Compliance**: GenerateSlug function now matches README documentation exactly
- **Case Sensitivity Control**: Preserve original case with `lowercase` parameter
- **Truncation Support**: Limit slug length with `maxLength` parameter
- **Enhanced Multi-language Support**: Improved diacritic removal for international characters

## 🔧 Improvements
- Variadic parameters for flexible API usage
- Comprehensive test coverage for all new features
- Example usage file demonstrating all functionality
- Complete API documentation

## 🔄 Backward Compatibility
- Original Slugger struct and methods preserved
- Existing code continues to work without changes

## 📝 Usage Examples
```go
// Basic usage
slug := slugger.GenerateSlug("Hello World!")
// Output: "hello-world"

// Case preservation
slug := slugger.GenerateSlug("Hello World!", "-", false)
// Output: "Hello-World"

// Truncation
slug := slugger.GenerateSlug("Very Long Title", "-", true, 10)
// Output: "very-long"
```

## 🚀 What's Next
Visit the [GitHub Releases page](https://github.com/jmrashed/go-slugger/releases) to create the official release.