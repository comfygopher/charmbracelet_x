// Package termspec provides terminal-specific specifications and detection.
// Each terminal specification contains methods for detecting if the current
// runtime is using that terminal and for calculating character widths
// specific to that terminal.
package termspec

import (
	"sync"

	"github.com/rivo/uniseg"
)

// Specification represents a terminal specification with detection and
// width calculation capabilities.
type Specification interface {
	// Detect returns true if the current runtime is using this terminal.
	Detect() bool

	// ColCount returns the number of columns that the given rune takes
	// when displayed in this terminal.
	ColCount(r rune) int

	// Name returns the name of this terminal specification.
	Name() string
}

var (
	// registry holds all registered terminal specifications
	registry = make([]Specification, 0)

	// registryMu protects the registry from concurrent access
	registryMu sync.RWMutex

	// current holds the detected current terminal specification
	current Specification

	// currentOnce ensures current terminal is detected only once
	currentOnce sync.Once
)

// Register adds a terminal specification to the registry.
// This is typically called from init() functions of terminal specifications.
func Register(spec Specification) {
	registryMu.Lock()
	defer registryMu.Unlock()
	registry = append(registry, spec)
}

// DetectCurrent detects and returns the current terminal specification.
// It iterates through all registered specifications and returns the first
// one that reports a positive detection. If no specific terminal is detected,
// it returns the default specification.
func DetectCurrent() Specification {
	currentOnce.Do(func() {
		registryMu.RLock()
		defer registryMu.RUnlock()

		// Try each registered specification
		for _, spec := range registry {
			if spec.Detect() {
				current = spec
				return
			}
		}

		// Fallback to default if no specific terminal detected
		current = &DefaultSpec{}
	})

	return current
}

// GetRegistered returns a copy of all registered terminal specifications.
func GetRegistered() []Specification {
	registryMu.RLock()
	defer registryMu.RUnlock()

	specs := make([]Specification, len(registry))
	copy(specs, registry)
	return specs
}

// DefaultSpec is the default terminal specification that uses standard
// width calculation methods. It serves as a fallback when no specific
// terminal is detected.
type DefaultSpec struct{}

// Name returns the name of the default specification.
func (d *DefaultSpec) Name() string {
	return "default"
}

// Detect always returns false for the default spec since it's used as fallback.
func (d *DefaultSpec) Detect() bool {
	return false
}

// ColCount returns the number of columns for a rune using the standard
// grapheme width calculation.
func (d *DefaultSpec) ColCount(r rune) int {
	// Use grapheme clustering by default for better Unicode support
	cluster := string(r)
	_, _, width, _ := uniseg.FirstGraphemeClusterInString(cluster, -1)
	return width
}
