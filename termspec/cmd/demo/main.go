package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/x/termspec"
)

func main() {
	fmt.Println("Terminal Specification Demo")
	fmt.Println("===========================")

	// Detect current terminal
	current := termspec.DetectCurrent()
	fmt.Printf("Current terminal: %s\n\n", current.Name())

	// Show all registered terminals
	fmt.Println("Registered terminal specifications:")
	for _, spec := range termspec.GetRegistered() {
		detected := ""
		if spec.Detect() {
			detected = " (DETECTED)"
		}
		fmt.Printf("- %s%s\n", spec.Name(), detected)
	}
	fmt.Println()

	// Test character width calculation
	fmt.Println("Character width calculation:")
	testChars := []rune{
		'A',  // ASCII letter
		'1',  // ASCII digit
		' ',  // Space
		'中',  // Chinese character
		'🚀',  // Rocket emoji
		'🌟',  // Star emoji
		'\t', // Tab
		'\n', // Newline
	}

	for _, char := range testChars {
		width := current.ColCount(char)
		if char == '\t' {
			fmt.Printf("'\\t' (tab) -> %d columns\n", width)
		} else if char == '\n' {
			fmt.Printf("'\\n' (newline) -> %d columns\n", width)
		} else {
			fmt.Printf("'%c' -> %d columns\n", char, width)
		}
	}
	fmt.Println()

	// Calculate string width
	testStrings := []string{
		"Hello",
		"Hello, World!",
		"Hello, 世界!",
		"🚀 Rocket Ship 🌟",
		"Café",
		"naïve résumé",
	}

	fmt.Println("String width calculation:")
	for _, str := range testStrings {
		width := stringWidth(current, str)
		fmt.Printf("'%s' -> %d columns\n", str, width)
	}
	fmt.Println()

	// Show environment variables used for detection
	fmt.Println("Environment variables used for terminal detection:")
	envVars := []string{
		"TERM",
		"TERM_PROGRAM",
		"TERMINAL_EMULATOR",
		"KITTY_WINDOW_ID",
		"KITTY_PID",
		"ITERM_SESSION_ID",
		"ITERM_PROFILE",
		"ALACRITTY_SOCKET",
		"WEZTERM_EXECUTABLE",
		"TMUX",
		"TMUX_PANE",
		"XTERM_VERSION",
	}

	for _, envVar := range envVars {
		value := os.Getenv(envVar)
		if value != "" {
			fmt.Printf("- %s=%s\n", envVar, value)
		} else {
			fmt.Printf("- %s=(not set)\n", envVar)
		}
	}
}

// stringWidth calculates the total width of a string using the given terminal specification
func stringWidth(spec termspec.Specification, s string) int {
	width := 0
	for _, r := range s {
		width += spec.ColCount(r)
	}
	return width
}
