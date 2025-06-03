package termspec

import (
	"os"
	"strings"

	"github.com/rivo/uniseg"
)

// WezTermSpec represents the WezTerm terminal specification.
type WezTermSpec struct{}

func init() {
	Register(&WezTermSpec{})
}

// Name returns the name of the WezTerm terminal specification.
func (w *WezTermSpec) Name() string {
	return "wezterm"
}

// Detect returns true if the current runtime is using WezTerm terminal.
func (w *WezTermSpec) Detect() bool {
	// Check TERM_PROGRAM environment variable
	termProgram := os.Getenv("TERM_PROGRAM")
	if termProgram == "WezTerm" {
		return true
	}

	// Check TERM environment variable
	term := os.Getenv("TERM")
	if strings.HasPrefix(term, "wezterm") {
		return true
	}

	// Check for WezTerm-specific environment variables
	if os.Getenv("WEZTERM_EXECUTABLE") != "" {
		return true
	}

	if os.Getenv("WEZTERM_CONFIG_FILE") != "" {
		return true
	}

	if os.Getenv("WEZTERM_PANE") != "" {
		return true
	}

	return false
}

// ColCount returns the number of columns that the given rune takes
// when displayed in WezTerm. WezTerm has excellent Unicode support
// and uses grapheme clustering.
func (w *WezTermSpec) ColCount(r rune) int {
	// WezTerm supports full grapheme clustering and has excellent Unicode support
	cluster := string(r)
	_, _, width, _ := uniseg.FirstGraphemeClusterInString(cluster, -1)
	return width
}
