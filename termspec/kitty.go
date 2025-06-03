package termspec

import (
	"os"
	"strings"

	"github.com/rivo/uniseg"
)

// KittySpec represents the Kitty terminal specification.
type KittySpec struct{}

func init() {
	Register(&KittySpec{})
}

// Name returns the name of the Kitty terminal specification.
func (k *KittySpec) Name() string {
	return "kitty"
}

// Detect returns true if the current runtime is using Kitty terminal.
// Kitty sets several environment variables that can be used for detection.
func (k *KittySpec) Detect() bool {
	// Check for Kitty-specific environment variables
	if os.Getenv("KITTY_WINDOW_ID") != "" {
		return true
	}

	if os.Getenv("KITTY_PID") != "" {
		return true
	}

	// Check TERM environment variable
	term := os.Getenv("TERM")
	if strings.HasPrefix(term, "xterm-kitty") {
		return true
	}

	// Check TERM_PROGRAM
	termProgram := os.Getenv("TERM_PROGRAM")
	if termProgram == "kitty" {
		return true
	}

	return false
}

// ColCount returns the number of columns that the given rune takes
// when displayed in Kitty terminal. Kitty has excellent Unicode support
// and uses grapheme clustering.
func (k *KittySpec) ColCount(r rune) int {
	// Kitty supports full grapheme clustering and has excellent Unicode support
	cluster := string(r)
	_, _, width, _ := uniseg.FirstGraphemeClusterInString(cluster, -1)
	return width
}
