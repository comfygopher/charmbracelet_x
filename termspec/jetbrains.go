package termspec

import (
	"os"

	"github.com/mattn/go-runewidth"
)

// JetBrainsSpec represents the JetBrains terminal specification.
type JetBrainsSpec struct{}

func init() {
	Register(&JetBrainsSpec{})
}

// Name returns the name of the JetBrains terminal specification.
func (j *JetBrainsSpec) Name() string {
	return "jetbrains"
}

// Detect returns true if the current runtime is using JetBrains terminal.
func (j *JetBrainsSpec) Detect() bool {
	// Check for JetBrains-specific environment variable
	terminalEmulator := os.Getenv("TERMINAL_EMULATOR")
	if terminalEmulator == "JetBrains-JediTerm" {
		return true
	}

	return false
}

// ColCount returns the number of columns that the given rune takes
// when displayed in JetBrains terminal. JetBrains terminal uses
// wcwidth-style calculation.
func (j *JetBrainsSpec) ColCount(r rune) int {
	// JetBrains terminal uses wcwidth-style calculation
	return runewidth.RuneWidth(r)
}
