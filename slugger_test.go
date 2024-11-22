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

	// Test case 4: Slugify with custom delimiter and non-alphabetic characters
	slugger = NewSlugger("_", true)
	result = slugger.Slugify("My--Custom--Slugger!2024")
	expected = "my_custom_slugger_2024"
	assert.Equal(t, expected, result)
}
