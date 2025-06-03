package termspec

import (
	"os"
	"strings"

	"github.com/rivo/uniseg"
)

// ITerm2Spec represents the iTerm2 terminal specification.
type ITerm2Spec struct{}

func init() {
	Register(&ITerm2Spec{})
}

// Name returns the name of the iTerm2 terminal specification.
func (i *ITerm2Spec) Name() string {
	return "iterm2"
}

// Detect returns true if the current runtime is using iTerm2 terminal.
func (i *ITerm2Spec) Detect() bool {
	// Check TERM_PROGRAM environment variable
	termProgram := os.Getenv("TERM_PROGRAM")
	if termProgram == "iTerm.app" {
		return true
	}

	// Check for iTerm2-specific environment variables
	if os.Getenv("ITERM_SESSION_ID") != "" {
		return true
	}

	if os.Getenv("ITERM_PROFILE") != "" {
		return true
	}

	// Check TERM environment variable
	term := os.Getenv("TERM")
	if strings.Contains(term, "iterm") {
		return true
	}

	return false
}

// ColCount returns the number of columns that the given rune takes
// when displayed in iTerm2. iTerm2 has good Unicode support.
func (i *ITerm2Spec) ColCount(r rune) int {
	// iTerm2 supports grapheme clustering and has good Unicode support
	cluster := string(r)
	_, _, width, _ := uniseg.FirstGraphemeClusterInString(cluster, -1)
	return width
}
