package slugger

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// Slugger struct to hold customization options.
type Slugger struct {
	Delimiter        string
	RemoveDiacritics bool
}

// NewSlugger creates a new instance of the Slugger.
func GenerateSlug(delimiter string, removeDiacritics bool) *Slugger {
	return &Slugger{
		Delimiter:        delimiter,
		RemoveDiacritics: removeDiacritics,
	}
}

// Slugify converts input string to a slug with customization options.
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
