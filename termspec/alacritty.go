package termspec

import (
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
)

// AlacrittySpec represents the Alacritty terminal specification.
type AlacrittySpec struct{}

func init() {
	Register(&AlacrittySpec{})
}

// Name returns the name of the Alacritty terminal specification.
func (a *AlacrittySpec) Name() string {
	return "alacritty"
}

// Detect returns true if the current runtime is using Alacritty terminal.
func (a *AlacrittySpec) Detect() bool {
	// Check TERM_PROGRAM environment variable
	termProgram := os.Getenv("TERM_PROGRAM")
	if termProgram == "alacritty" {
		return true
	}

	// Check TERM environment variable
	term := os.Getenv("TERM")
	if strings.HasPrefix(term, "alacritty") {
		return true
	}

	// Check for Alacritty-specific environment variables
	if os.Getenv("ALACRITTY_SOCKET") != "" {
		return true
	}

	if os.Getenv("ALACRITTY_LOG") != "" {
		return true
	}

	return false
}

// ColCount returns the number of columns that the given rune takes
// when displayed in Alacritty. Alacritty uses wcwidth-style calculation.
func (a *AlacrittySpec) ColCount(r rune) int {
	// Alacritty traditionally uses wcwidth-style calculation
	// rather than full grapheme clustering
	return runewidth.RuneWidth(r)
}
