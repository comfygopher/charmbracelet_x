package termspec

import (
	"os"
	"strings"

	"github.com/mattn/go-runewidth"
)

// TmuxSpec represents the Tmux terminal specification.
type TmuxSpec struct{}

func init() {
	Register(&TmuxSpec{})
}

// Name returns the name of the Tmux terminal specification.
func (t *TmuxSpec) Name() string {
	return "tmux"
}

// Detect returns true if the current runtime is using Tmux.
func (t *TmuxSpec) Detect() bool {
	// Check for Tmux-specific environment variables
	if os.Getenv("TMUX") != "" {
		return true
	}

	if os.Getenv("TMUX_PANE") != "" {
		return true
	}

	// Check TERM environment variable
	term := os.Getenv("TERM")
	if strings.HasPrefix(term, "tmux") {
		return true
	}

	if strings.HasPrefix(term, "screen") && os.Getenv("TMUX") != "" {
		return true
	}

	return false
}

// ColCount returns the number of columns that the given rune takes
// when displayed in Tmux. Tmux generally uses wcwidth calculation.
func (t *TmuxSpec) ColCount(r rune) int {
	// Tmux uses wcwidth-style calculation
	return runewidth.RuneWidth(r)
}
