package main

import (
	"fmt"
	"github.com/jmrashed/go-slugger/slugger"
)

func main() {
	// Basic usage
	slug1 := slugger.GenerateSlug("Hello World! Welcome to Go.")
	fmt.Printf("Basic: %s\n", slug1) // Output: hello-world-welcome-to-go

	// Custom delimiter
	slug2 := slugger.GenerateSlug("Hello World!", "_")
	fmt.Printf("Custom delimiter: %s\n", slug2) // Output: hello_world

	// Diacritic removal
	slug3 := slugger.GenerateSlug("Café de Paris")
	fmt.Printf("Diacritic removal: %s\n", slug3) // Output: cafe-de-paris

	// Case preservation
	slug4 := slugger.GenerateSlug("Hello World!", "-", false)
	fmt.Printf("Case preservation: %s\n", slug4) // Output: Hello-World

	// Truncation
	slug5 := slugger.GenerateSlug("This is a very long title", "-", true, 20)
	fmt.Printf("Truncation: %s\n", slug5) // Output: this-is-a-very-lo

	// Multi-language support (German)
	slug6 := slugger.GenerateSlug("München ist schön!")
	fmt.Printf("German: %s\n", slug6) // Output: munchen-ist-schon

	// Case preservation with custom delimiter
	slug7 := slugger.GenerateSlug("München ist schön!", "_", false)
	fmt.Printf("German with underscore: %s\n", slug7) // Output: Munchen_ist_schon

	// Truncation with custom length
	slug8 := slugger.GenerateSlug("München ist schön!", "-", true, 15)
	fmt.Printf("German truncated: %s\n", slug8) // Output: munchen-ist-sch
}