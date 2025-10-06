package slugger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test new GenerateSlug API
func TestGenerateSlug(t *testing.T) {
	// Test case 1: Basic usage with default settings
	result := GenerateSlug("Hello World! Welcome to Go.")
	expected := "hello-world-welcome-to-go"
	assert.Equal(t, expected, result)

	// Test case 2: Custom delimiter
	result = GenerateSlug("Hello World!", "_")
	expected = "hello_world"
	assert.Equal(t, expected, result)

	// Test case 3: Diacritic removal
	result = GenerateSlug("Café de Paris")
	expected = "cafe-de-paris"
	assert.Equal(t, expected, result)

	// Test case 4: Case preservation
	result = GenerateSlug("Hello World!", "-", false)
	expected = "Hello-World"
	assert.Equal(t, expected, result)

	// Test case 5: Truncation
	result = GenerateSlug("This is a very long title", "-", true, 20)
	expected = "this-is-a-very-lo"
	assert.Equal(t, expected, result)

	// Test case 6: Multi-language support (German)
	result = GenerateSlug("München ist schön!")
	expected = "munchen-ist-schon"
	assert.Equal(t, expected, result)

	// Test case 7: Case preservation with custom delimiter
	result = GenerateSlug("München ist schön!", "_", false)
	expected = "Munchen_ist_schon"
	assert.Equal(t, expected, result)

	// Test case 8: Truncation with custom length
	result = GenerateSlug("München ist schön!", "-", true, 15)
	expected = "munchen-ist-sch"
	assert.Equal(t, expected, result)
}

// Test backward compatibility with old Slugger API
func TestSlugger(t *testing.T) {
	// Test case 1: Basic Slugify with default settings
	slugger := NewSlugger("-", true)
	result := slugger.Slugify("Hello World! This is a test.")
	expected := "hello-world-this-is-a-test"
	assert.Equal(t, expected, result)

	// Test case 2: Slugify without diacritic removal
	slugger = NewSlugger("-", false)
	result = slugger.Slugify("Café au lait")
	expected = "cafe-au-lait"
	assert.Equal(t, expected, result)

	// Test case 3: Slugify with custom delimiter
	slugger = NewSlugger("_", true)
	result = slugger.Slugify("Hello World! This is a test.")
	expected = "hello_world_this_is_a_test"
	assert.Equal(t, expected, result)

	// Test case 4: Slugify with custom delimiter and non-alphabetic characters
	slugger = NewSlugger("_", true)
	result = slugger.Slugify("My--Custom--Slugger!2024")
	expected = "my_custom_slugger_2024"
	assert.Equal(t, expected, result)
}
