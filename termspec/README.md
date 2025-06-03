# termspec

Package `termspec` provides terminal-specific specifications and detection for Go applications. Each terminal specification contains methods for detecting if the current runtime is using that terminal and for calculating character widths specific to that terminal.

## Features

- **Automatic Terminal Detection**: Detects the current terminal based on environment variables
- **Terminal-Specific Width Calculation**: Each terminal has its own method for calculating character column widths
- **Extensible Architecture**: Easy to add new terminal specifications
- **Thread-Safe**: Safe for concurrent use
- **Fallback Support**: Always provides a default specification if no specific terminal is detected

## Supported Terminals

- **Kitty**: Advanced terminal with excellent Unicode support
- **iTerm2**: Popular macOS terminal with good Unicode support
- **Alacritty**: GPU-accelerated terminal using wcwidth calculation
- **WezTerm**: Modern terminal with full grapheme clustering support
- **XTerm**: Traditional terminal using wcwidth calculation
- **Tmux**: Terminal multiplexer using wcwidth calculation
- **JetBrains**: JetBrains IDE terminal using wcwidth calculation
- **Default**: Fallback specification using grapheme clustering

## Installation

```bash
go get github.com/charmbracelet/x/termspec
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/charmbracelet/x/termspec"
)

func main() {
    // Detect current terminal
    spec := termspec.DetectCurrent()
    fmt.Printf("Current terminal: %s\n", spec.Name())
    
    // Calculate character widths
    fmt.Printf("'A' takes %d columns\n", spec.ColCount('A'))
    fmt.Printf("'中' takes %d columns\n", spec.ColCount('中'))
    fmt.Printf("'🚀' takes %d columns\n", spec.ColCount('🚀'))
}
```

## API Reference

### Core Functions

#### `DetectCurrent() Specification`
Detects and returns the current terminal specification. This function is thread-safe and caches the result.

#### `GetRegistered() []Specification`
Returns a copy of all registered terminal specifications.

#### `Register(spec Specification)`
Registers a new terminal specification. Typically called from `init()` functions.

### Specification Interface

```go
type Specification interface {
    // Detect returns true if the current runtime is using this terminal
    Detect() bool
    
    // ColCount returns the number of columns that the given rune takes
    // when displayed in this terminal
    ColCount(r rune) int
    
    // Name returns the name of this terminal specification
    Name() string
}
```

## Detection Logic

The package detects terminals using environment variables:

| Terminal | Environment Variables |
|----------|----------------------|
| Kitty | `KITTY_WINDOW_ID`, `KITTY_PID`, `TERM=xterm-kitty*`, `TERM_PROGRAM=kitty` |
| iTerm2 | `TERM_PROGRAM=iTerm.app`, `ITERM_SESSION_ID`, `ITERM_PROFILE` |
| Alacritty | `TERM_PROGRAM=alacritty`, `TERM=alacritty*`, `ALACRITTY_SOCKET` |
| WezTerm | `TERM_PROGRAM=WezTerm`, `TERM=wezterm*`, `WEZTERM_EXECUTABLE` |
| XTerm | `TERM=xterm*` (when no other terminal detected), `XTERM_VERSION` |
| Tmux | `TMUX`, `TMUX_PANE`, `TERM=tmux*` |
| JetBrains | `TERMINAL_EMULATOR=JetBrains-JediTerm` |

## Width Calculation Methods

Different terminals use different methods for calculating character widths:

- **Grapheme Clustering**: Used by Kitty, iTerm2, WezTerm, and Default
  - Provides better Unicode support
  - Handles complex characters and emoji properly
  
- **wcwidth**: Used by Alacritty, XTerm, Tmux, and JetBrains
  - Traditional width calculation
  - Faster but less accurate for complex Unicode

## Creating Custom Specifications

You can create custom terminal specifications by implementing the `Specification` interface:

```go
type MyTerminalSpec struct{}

func (m *MyTerminalSpec) Name() string {
    return "myterminal"
}

func (m *MyTerminalSpec) Detect() bool {
    return os.Getenv("MY_TERMINAL") != ""
}

func (m *MyTerminalSpec) ColCount(r rune) int {
    // Custom width calculation logic
    if r < 128 {
        return 1 // ASCII
    }
    return 2 // Non-ASCII
}

func init() {
    termspec.Register(&MyTerminalSpec{})
}
```

## Examples

### String Width Calculation

```go
func stringWidth(s string) int {
    spec := termspec.DetectCurrent()
    width := 0
    for _, r := range s {
        width += spec.ColCount(r)
    }
    return width
}

width := stringWidth("Hello, 世界! 🚀")
fmt.Printf("String width: %d columns\n", width)
```

### Terminal-Specific Behavior

```go
// Check if running in a specific terminal
if spec := termspec.DetectCurrent(); spec.Name() == "kitty" {
    fmt.Println("Running in Kitty - full Unicode support available")
}

// Compare width calculation across terminals
terminals := []termspec.Specification{
    &termspec.KittySpec{},
    &termspec.AlacrittySpec{},
}

for _, term := range terminals {
    width := term.ColCount('🚀')
    fmt.Printf("%s: emoji width = %d\n", term.Name(), width)
}
```

## Thread Safety

All functions in this package are thread-safe. The terminal detection is performed only once and cached for subsequent calls.

## Performance

- Terminal detection: O(1) after first call (cached)
- Width calculation: O(1) per character
- Registration: O(1) per specification

## Contributing

To add support for a new terminal:

1. Create a new file (e.g., `newterminal.go`)
2. Implement the `Specification` interface
3. Add detection logic based on environment variables
4. Register the specification in an `init()` function
5. Add tests in `termspec_test.go`

## License

This package is part of the Charm Bracelet X collection and follows the same license terms.
