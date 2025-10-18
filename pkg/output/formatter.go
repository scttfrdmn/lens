package output

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Color constants using ANSI escape codes
const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
	colorBold   = "\033[1m"
)

// Formatter handles colored and formatted output
type Formatter struct {
	writer   io.Writer
	noColor  bool
	useEmoji bool
}

// NewFormatter creates a new Formatter instance
func NewFormatter(w io.Writer, noColor bool) *Formatter {
	return &Formatter{
		writer:   w,
		noColor:  noColor,
		useEmoji: true, // Enable emoji by default
	}
}

// DefaultFormatter returns a formatter writing to stdout
func DefaultFormatter() *Formatter {
	return NewFormatter(os.Stdout, false)
}

// colorize wraps text in color codes if colors are enabled
func (f *Formatter) colorize(color, text string) string {
	if f.noColor {
		return text
	}
	return color + text + colorReset
}

// Success prints a success message in green with checkmark
func (f *Formatter) Success(message string) {
	emoji := "✓"
	if !f.useEmoji {
		emoji = "[OK]"
	}
	fmt.Fprintf(f.writer, "%s %s\n", f.colorize(colorGreen, emoji), message)
}

// SuccessWithDetail prints a success message with additional detail
func (f *Formatter) SuccessWithDetail(message, detail string) {
	emoji := "✓"
	if !f.useEmoji {
		emoji = "[OK]"
	}
	fmt.Fprintf(f.writer, "%s %s: %s\n", f.colorize(colorGreen, emoji), message, detail)
}

// Error prints an error message in red with X mark
func (f *Formatter) Error(message string) {
	emoji := "✗"
	if !f.useEmoji {
		emoji = "[ERROR]"
	}
	fmt.Fprintf(f.writer, "%s %s\n", f.colorize(colorRed, emoji), message)
}

// Warning prints a warning message in yellow with warning symbol
func (f *Formatter) Warning(message string) {
	emoji := "⚠️"
	if !f.useEmoji {
		emoji = "[WARN]"
	}
	fmt.Fprintf(f.writer, "%s  %s\n", emoji, f.colorize(colorYellow, message))
}

// Info prints an informational message in blue with info symbol
func (f *Formatter) Info(message string) {
	emoji := "ℹ️"
	if !f.useEmoji {
		emoji = "[INFO]"
	}
	fmt.Fprintf(f.writer, "%s  %s\n", emoji, message)
}

// Step prints a step message with an emoji indicator
func (f *Formatter) Step(emoji, message string) {
	if !f.useEmoji {
		emoji = ">"
	}
	fmt.Fprintf(f.writer, "%s %s\n", emoji, message)
}

// StepWithDetail prints a step message with additional detail
func (f *Formatter) StepWithDetail(emoji, message, detail string) {
	if !f.useEmoji {
		emoji = ">"
	}
	fmt.Fprintf(f.writer, "%s %s: %s\n", emoji, message, f.colorize(colorCyan, detail))
}

// Status prints a status update (typically during operations)
func (f *Formatter) Status(message string) {
	emoji := "⏳"
	if !f.useEmoji {
		emoji = "..."
	}
	fmt.Fprintf(f.writer, "%s %s\n", emoji, message)
}

// Progress prints indented progress messages (for sub-steps)
func (f *Formatter) Progress(message string) {
	emoji := "📋"
	if !f.useEmoji {
		emoji = "-"
	}
	fmt.Fprintf(f.writer, "   %s %s\n", emoji, message)
}

// ProgressWithTime prints progress with elapsed time
func (f *Formatter) ProgressWithTime(message string, elapsedSeconds int) {
	fmt.Fprintf(f.writer, "   [%ds] %s\n", elapsedSeconds, message)
}

// Complete prints a completion message with celebration emoji
func (f *Formatter) Complete(message string) {
	emoji := "🎉"
	if !f.useEmoji {
		emoji = "[DONE]"
	}
	fmt.Fprintf(f.writer, "%s %s\n", emoji, f.colorize(colorGreen+colorBold, message))
}

// Header prints a header message
func (f *Formatter) Header(message string) {
	fmt.Fprintf(f.writer, "\n%s\n", f.colorize(colorBold, message))
}

// Subheader prints a subheader message
func (f *Formatter) Subheader(message string) {
	fmt.Fprintf(f.writer, "\n%s\n", message)
}

// KeyValue prints a key-value pair with formatting
func (f *Formatter) KeyValue(key, value string) {
	fmt.Fprintf(f.writer, "  %s: %s\n", f.colorize(colorBold, key), value)
}

// List prints a bulleted list item
func (f *Formatter) List(item string) {
	fmt.Fprintf(f.writer, "  - %s\n", item)
}

// DryRun prints a message with [DRY RUN] prefix
func (f *Formatter) DryRun(message string) {
	prefix := f.colorize(colorYellow+colorBold, "[DRY RUN]")
	fmt.Fprintf(f.writer, "%s %s\n", prefix, message)
}

// Cost prints a cost-related message with money emoji
func (f *Formatter) Cost(message string) {
	emoji := "💰"
	if !f.useEmoji {
		emoji = "[$]"
	}
	fmt.Fprintf(f.writer, "%s %s\n", emoji, f.colorize(colorYellow, message))
}

// Connection prints connection information with link emoji
func (f *Formatter) Connection(message string) {
	emoji := "🔗"
	if !f.useEmoji {
		emoji = "[LINK]"
	}
	fmt.Fprintf(f.writer, "%s %s\n", emoji, f.colorize(colorCyan, message))
}

// Separator prints a visual separator line
func (f *Formatter) Separator() {
	fmt.Fprintln(f.writer, strings.Repeat("─", 60))
}

// Blank prints a blank line
func (f *Formatter) Blank() {
	fmt.Fprintln(f.writer)
}

// Print prints a plain message without formatting
func (f *Formatter) Print(message string) {
	fmt.Fprintln(f.writer, message)
}

// Printf prints a formatted message without additional styling
func (f *Formatter) Printf(format string, args ...interface{}) {
	fmt.Fprintf(f.writer, format, args...)
}
