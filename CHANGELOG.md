# Changelog

## [1.1.0] - 2024-11-22
### Added
- New `GenerateSlug` function with variadic parameters matching README API
- Case sensitivity control (preserve original case option)
- Slug truncation functionality with maxLength parameter
- Enhanced multi-language support with improved diacritic removal
- Comprehensive test coverage for all new features
- Example usage file demonstrating all functionality

### Changed
- API now matches README documentation exactly
- Improved regex pattern handling for case-sensitive slugs

### Backward Compatibility
- Original Slugger struct and methods preserved for backward compatibility

## [1.0.1] - 2024-11-22
### Added
- Initial version of go-slugger with support for slug generation and diacritic removal.
