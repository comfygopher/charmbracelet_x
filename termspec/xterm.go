package termspec

import (
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
)

// XTermSpec represents the XTerm terminal specification.
type XTermSpec struct{}

func init() {
	Register(&XTermSpec{})
}

// Name returns the name of the XTerm terminal specification.
func (x *XTermSpec) Name() string {
	return "xterm"
}

// Detect returns true if the current runtime is using XTerm terminal.
func (x *XTermSpec) Detect() bool {
	// Check TERM environment variable
	term := os.Getenv("TERM")
	if strings.HasPrefix(term, "xterm") {
		// Make sure it's not a more specific terminal that uses xterm as base
		termProgram := os.Getenv("TERM_PROGRAM")
		if termProgram == "" || termProgram == "xterm" {
			return true
		}
	}

	// Check for XTerm-specific environment variables
	if os.Getenv("XTERM_VERSION") != "" {
		return true
	}

	return false
}

// ColCount returns the number of columns that the given rune takes
// when displayed in XTerm. XTerm uses traditional wcwidth calculation.
func (x *XTermSpec) ColCount(r rune) int {
	// XTerm uses traditional wcwidth-style calculation
	return runewidth.RuneWidth(r)
}
