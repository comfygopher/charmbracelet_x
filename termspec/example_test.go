package termspec_test

import (
	"fmt"
	"log"

	"github.com/charmbracelet/x/termspec"
)

func ExampleDetectCurrent() {
	// Detect the current terminal specification
	spec := termspec.DetectCurrent()

	fmt.Printf("Current terminal: %s\n", spec.Name())

	// Calculate column width for various characters
	testRunes := []rune{'a', '中', '🚀', '\t'}
	for _, r := range testRunes {
		width := spec.ColCount(r)
		fmt.Printf("'%c' takes %d columns\n", r, width)
	}
}

func ExampleGetRegistered() {
	// Get all registered terminal specifications
	specs := termspec.GetRegistered()

	fmt.Printf("Available terminal specifications:\n")
	for _, spec := range specs {
		fmt.Printf("- %s\n", spec.Name())
	}
}

func ExampleSpecification() {
	// Example of using a specific terminal specification
	kitty := &termspec.KittySpec{}

	if kitty.Detect() {
		fmt.Println("Running in Kitty terminal")
	} else {
		fmt.Println("Not running in Kitty terminal")
	}

	// Calculate width for a wide character
	width := kitty.ColCount('中')
	fmt.Printf("Chinese character '中' takes %d columns in Kitty\n", width)
}

// CustomSpec is an example custom terminal specification
type CustomSpec struct{}

func (c *CustomSpec) Name() string {
	return "custom"
}

func (c *CustomSpec) Detect() bool {
	// Custom detection logic
	return false
}

func (c *CustomSpec) ColCount(r rune) int {
	// Custom width calculation
	if r < 128 {
		return 1 // ASCII characters
	}
	return 2 // Non-ASCII characters
}

func Example_customSpecification() {
	// Example of creating a custom terminal specification
	custom := &CustomSpec{}
	termspec.Register(custom)

	fmt.Printf("Registered custom specification: %s\n", custom.Name())
}

// Example showing how to use the package in a real application
func Example_usageInApplication() {
	// Detect current terminal
	spec := termspec.DetectCurrent()

	// Function to calculate string width using terminal-specific logic
	stringWidth := func(s string) int {
		width := 0
		for _, r := range s {
			width += spec.ColCount(r)
		}
		return width
	}

	// Example strings with different character types
	examples := []string{
		"Hello",     // ASCII
		"Hello, 世界", // Mixed ASCII and wide characters
		"🚀 Rocket",  // Emoji and ASCII
		"Café",      // Accented characters
	}

	fmt.Printf("String width calculation using %s terminal:\n", spec.Name())
	for _, example := range examples {
		width := stringWidth(example)
		fmt.Printf("'%s' -> %d columns\n", example, width)
	}
}

// Example showing terminal-specific behavior differences
func Example_terminalDifferences() {
	// Create instances of different terminal specifications
	terminals := []termspec.Specification{
		&termspec.KittySpec{},
		&termspec.AlacrittySpec{},
		&termspec.XTermSpec{},
		&termspec.WezTermSpec{},
		&termspec.JetBrainsSpec{},
	}

	// Test character that might have different widths
	testChar := '🚀' // Emoji

	fmt.Printf("Width of '%c' in different terminals:\n", testChar)
	for _, term := range terminals {
		width := term.ColCount(testChar)
		fmt.Printf("- %s: %d columns\n", term.Name(), width)
	}
}

// Example showing how to handle terminal detection errors gracefully
func Example_gracefulFallback() {
	// Always get a valid specification (falls back to default if needed)
	spec := termspec.DetectCurrent()

	// Log the detected terminal for debugging
	log.Printf("Using terminal specification: %s", spec.Name())

	// Use the specification safely
	width := spec.ColCount('A')
	fmt.Printf("Character 'A' width: %d\n", width)
}
