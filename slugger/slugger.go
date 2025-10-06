package slugger

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// GenerateSlug generates a slug from the input string with optional parameters
func GenerateSlug(input string, args ...interface{}) string {
	delimiter := "-"
	lowercase := true
	maxLength := 0

	// Parse optional arguments
	for i, arg := range args {
		switch i {
		case 0:
			if s, ok := arg.(string); ok {
				delimiter = s
			}
		case 1:
			if b, ok := arg.(bool); ok {
				lowercase = b
			}
		case 2:
			if i, ok := arg.(int); ok {
				maxLength = i
			}
		}
	}

	// Normalize string (decompose accents into separate characters)
	input = norm.NFD.String(input)
	input = strings.Map(removeDiacritics, input)

	// Convert to lowercase if specified
	if lowercase {
		input = strings.ToLower(input)
	}

	// Replace non-alphanumeric characters with delimiters
	pattern := `[^a-zA-Z0-9]+`
	if lowercase {
		pattern = `[^a-z0-9]+`
	}
	input = regexp.MustCompile(pattern).ReplaceAllString(input, delimiter)

	// Trim delimiters from the start and end
	input = strings.Trim(input, delimiter)

	// Truncate if maxLength is specified
	if maxLength > 0 && len(input) > maxLength {
		input = input[:maxLength]
		// Ensure we don't end with a delimiter after truncation
		input = strings.TrimRight(input, delimiter)
	}

	return input
}

// Slugger struct to hold customization options (kept for backward compatibility)
type Slugger struct {
	Delimiter        string
	RemoveDiacritics bool
}

// NewSlugger creates a new instance of the Slugger (kept for backward compatibility)
func NewSlugger(delimiter string, removeDiacritics bool) *Slugger {
	return &Slugger{
		Delimiter:        delimiter,
		RemoveDiacritics: removeDiacritics,
	}
}

// Slugify converts input string to a slug with customization options (kept for backward compatibility)
func (s *Slugger) Slugify(input string) string {
	// Normalize string (decompose accents into separate characters)
	if s.RemoveDiacritics {
		input = norm.NFD.String(input)
		input = strings.Map(removeDiacritics, input)
	}

	// Convert to lowercase
	input = strings.ToLower(input)

	// Replace non-alphanumeric characters with delimiters
	input = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(input, s.Delimiter)

	// Trim delimiters from the start and end
	input = strings.Trim(input, s.Delimiter)

	return input
}

// removeDiacritics removes diacritical marks (accents) from a rune.
func removeDiacritics(r rune) rune {
	// Keep the characters that don't have accents or remove accent marks
	if unicode.IsMark(r) {
		return -1
	}
	return r
}
