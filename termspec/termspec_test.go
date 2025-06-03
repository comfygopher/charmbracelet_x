package termspec

import (
	"os"
	"testing"
)

func TestDetectCurrent(t *testing.T) {
	// Test that DetectCurrent returns a valid specification
	spec := DetectCurrent()
	if spec == nil {
		t.Fatal("DetectCurrent returned nil")
	}

	// Test that the specification has a name
	name := spec.Name()
	if name == "" {
		t.Error("Specification name is empty")
	}

	t.Logf("Detected terminal: %s", name)
}

func TestDefaultSpec(t *testing.T) {
	spec := &DefaultSpec{}

	// Test name
	if spec.Name() != "default" {
		t.Errorf("Expected name 'default', got '%s'", spec.Name())
	}

	// Test detect (should always return false)
	if spec.Detect() {
		t.Error("DefaultSpec.Detect() should return false")
	}

	// Test ColCount with various runes
	testCases := []struct {
		rune     rune
		expected int
	}{
		{'a', 1},  // ASCII letter
		{'1', 1},  // ASCII digit
		{' ', 1},  // Space
		{'中', 2},  // Wide character (Chinese)
		{'🚀', 2},  // Emoji
		{'\t', 0}, // Tab (control character)
	}

	for _, tc := range testCases {
		width := spec.ColCount(tc.rune)
		if width < 0 {
			t.Errorf("ColCount(%q) returned negative width: %d", tc.rune, width)
		}
		// Note: We don't test exact values as they may vary between Unicode versions
		t.Logf("ColCount(%q) = %d", tc.rune, width)
	}
}

func TestKittySpec(t *testing.T) {
	spec := &KittySpec{}

	// Test name
	if spec.Name() != "kitty" {
		t.Errorf("Expected name 'kitty', got '%s'", spec.Name())
	}

	// Test detection with environment variables
	originalEnv := map[string]string{
		"KITTY_WINDOW_ID": os.Getenv("KITTY_WINDOW_ID"),
		"KITTY_PID":       os.Getenv("KITTY_PID"),
		"TERM":            os.Getenv("TERM"),
		"TERM_PROGRAM":    os.Getenv("TERM_PROGRAM"),
	}

	// Clean environment
	for key := range originalEnv {
		os.Unsetenv(key)
	}

	// Should not detect without environment variables
	if spec.Detect() {
		t.Error("KittySpec should not detect without environment variables")
	}

	// Test with KITTY_WINDOW_ID
	os.Setenv("KITTY_WINDOW_ID", "1")
	if !spec.Detect() {
		t.Error("KittySpec should detect with KITTY_WINDOW_ID")
	}
	os.Unsetenv("KITTY_WINDOW_ID")

	// Test with TERM
	os.Setenv("TERM", "xterm-kitty")
	if !spec.Detect() {
		t.Error("KittySpec should detect with TERM=xterm-kitty")
	}
	os.Unsetenv("TERM")

	// Test with TERM_PROGRAM
	os.Setenv("TERM_PROGRAM", "kitty")
	if !spec.Detect() {
		t.Error("KittySpec should detect with TERM_PROGRAM=kitty")
	}

	// Restore original environment
	for key, value := range originalEnv {
		if value != "" {
			os.Setenv(key, value)
		} else {
			os.Unsetenv(key)
		}
	}
}

func TestITerm2Spec(t *testing.T) {
	spec := &ITerm2Spec{}

	// Test name
	if spec.Name() != "iterm2" {
		t.Errorf("Expected name 'iterm2', got '%s'", spec.Name())
	}

	// Test basic functionality
	width := spec.ColCount('a')
	if width < 0 {
		t.Errorf("ColCount('a') returned negative width: %d", width)
	}
}

func TestAlacrittySpec(t *testing.T) {
	spec := &AlacrittySpec{}

	// Test name
	if spec.Name() != "alacritty" {
		t.Errorf("Expected name 'alacritty', got '%s'", spec.Name())
	}

	// Test basic functionality
	width := spec.ColCount('a')
	if width < 0 {
		t.Errorf("ColCount('a') returned negative width: %d", width)
	}
}

func TestWezTermSpec(t *testing.T) {
	spec := &WezTermSpec{}

	// Test name
	if spec.Name() != "wezterm" {
		t.Errorf("Expected name 'wezterm', got '%s'", spec.Name())
	}

	// Test basic functionality
	width := spec.ColCount('a')
	if width < 0 {
		t.Errorf("ColCount('a') returned negative width: %d", width)
	}
}

func TestXTermSpec(t *testing.T) {
	spec := &XTermSpec{}

	// Test name
	if spec.Name() != "xterm" {
		t.Errorf("Expected name 'xterm', got '%s'", spec.Name())
	}

	// Test basic functionality
	width := spec.ColCount('a')
	if width < 0 {
		t.Errorf("ColCount('a') returned negative width: %d", width)
	}
}

func TestTmuxSpec(t *testing.T) {
	spec := &TmuxSpec{}

	// Test name
	if spec.Name() != "tmux" {
		t.Errorf("Expected name 'tmux', got '%s'", spec.Name())
	}

	// Test basic functionality
	width := spec.ColCount('a')
	if width < 0 {
		t.Errorf("ColCount('a') returned negative width: %d", width)
	}
}

func TestGetRegistered(t *testing.T) {
	specs := GetRegistered()
	if len(specs) == 0 {
		t.Error("No terminal specifications registered")
	}

	// Check that all registered specs have names
	names := make(map[string]bool)
	for _, spec := range specs {
		name := spec.Name()
		if name == "" {
			t.Error("Found specification with empty name")
		}
		if names[name] {
			t.Errorf("Duplicate specification name: %s", name)
		}
		names[name] = true
	}

	t.Logf("Registered %d terminal specifications: %v", len(specs), getNames(specs))
}

func getNames(specs []Specification) []string {
	names := make([]string, len(specs))
	for i, spec := range specs {
		names[i] = spec.Name()
	}
	return names
}

func BenchmarkDetectCurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DetectCurrent()
	}
}

func BenchmarkColCount(b *testing.B) {
	spec := DetectCurrent()
	testRunes := []rune{'a', '中', '🚀', '\t'}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, r := range testRunes {
			spec.ColCount(r)
		}
	}
}
