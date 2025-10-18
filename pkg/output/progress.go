package output

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// ProgressBar displays a text-based progress bar
type ProgressBar struct {
	writer      io.Writer
	total       int
	current     int
	width       int
	description string
	startTime   time.Time
	noColor     bool
}

// NewProgressBar creates a new progress bar
func NewProgressBar(total int, description string) *ProgressBar {
	return &ProgressBar{
		writer:      os.Stdout,
		total:       total,
		current:     0,
		width:       20, // Default width of 20 characters
		description: description,
		startTime:   time.Now(),
		noColor:     false,
	}
}

// SetWidth sets the width of the progress bar
func (pb *ProgressBar) SetWidth(width int) {
	pb.width = width
}

// SetNoColor disables color output
func (pb *ProgressBar) SetNoColor(noColor bool) {
	pb.noColor = noColor
}

// Update updates the progress bar to a specific value
func (pb *ProgressBar) Update(current int) {
	pb.current = current
	pb.render()
}

// Increment increments the progress bar by one
func (pb *ProgressBar) Increment() {
	pb.current++
	pb.render()
}

// render draws the progress bar
func (pb *ProgressBar) render() {
	percentage := float64(pb.current) / float64(pb.total) * 100
	filled := int(float64(pb.width) * float64(pb.current) / float64(pb.total))

	// Build the bar
	var bar strings.Builder
	bar.WriteString("[")

	// Filled portion
	for i := 0; i < filled; i++ {
		bar.WriteString("█")
	}

	// Empty portion
	for i := filled; i < pb.width; i++ {
		bar.WriteString("░")
	}

	bar.WriteString("]")

	// Calculate elapsed time and ETA
	elapsed := time.Since(pb.startTime)
	var eta string
	if pb.current > 0 {
		totalEstimated := elapsed * time.Duration(pb.total) / time.Duration(pb.current)
		remaining := totalEstimated - elapsed
		eta = fmt.Sprintf(" (ETA: %s)", formatDuration(remaining))
	}

	// Print the progress bar (with carriage return to overwrite)
	barStr := bar.String()
	if !pb.noColor {
		barStr = colorGreen + barStr + colorReset
	}

	fmt.Fprintf(pb.writer, "\r%s %s %.0f%%%s", pb.description, barStr, percentage, eta)

	// If complete, add newline
	if pb.current >= pb.total {
		fmt.Fprintln(pb.writer)
	}
}

// Finish marks the progress bar as complete
func (pb *ProgressBar) Finish() {
	pb.current = pb.total
	pb.render()
}

// SimpleProgress represents a simple text-based progress indicator
type SimpleProgress struct {
	writer      io.Writer
	description string
	startTime   time.Time
	lastUpdate  time.Time
}

// NewSimpleProgress creates a simple progress indicator (no bar, just text updates)
func NewSimpleProgress(description string) *SimpleProgress {
	return &SimpleProgress{
		writer:      os.Stdout,
		description: description,
		startTime:   time.Now(),
		lastUpdate:  time.Now(),
	}
}

// Update prints a progress update with elapsed time
func (sp *SimpleProgress) Update(message string) {
	elapsed := time.Since(sp.startTime)
	fmt.Fprintf(sp.writer, "   [%s] %s\n", formatDuration(elapsed), message)
	sp.lastUpdate = time.Now()
}

// UpdateWithElapsed prints a progress update showing elapsed seconds
func (sp *SimpleProgress) UpdateWithElapsed(message string) {
	elapsed := int(time.Since(sp.startTime).Seconds())
	fmt.Fprintf(sp.writer, "   [%ds] %s\n", elapsed, message)
	sp.lastUpdate = time.Now()
}

// Finish marks the progress as complete
func (sp *SimpleProgress) Finish(message string) {
	elapsed := time.Since(sp.startTime)
	fmt.Fprintf(sp.writer, "   ✅ %s (completed in %s)\n", message, formatDuration(elapsed))
}

// formatDuration formats a duration in a human-readable way
func formatDuration(d time.Duration) string {
	if d < time.Second {
		return "< 1s"
	}

	seconds := int(d.Seconds())
	if seconds < 60 {
		return fmt.Sprintf("%ds", seconds)
	}

	minutes := seconds / 60
	seconds = seconds % 60

	if minutes < 60 {
		if seconds > 0 {
			return fmt.Sprintf("%dm%ds", minutes, seconds)
		}
		return fmt.Sprintf("%dm", minutes)
	}

	hours := minutes / 60
	minutes = minutes % 60

	if minutes > 0 {
		return fmt.Sprintf("%dh%dm", hours, minutes)
	}
	return fmt.Sprintf("%dh", hours)
}

// EstimatedTimeMessage returns a user-friendly estimated time message
func EstimatedTimeMessage(estimatedSeconds int) string {
	if estimatedSeconds < 60 {
		return fmt.Sprintf("this should take about %d seconds", estimatedSeconds)
	}

	minutes := estimatedSeconds / 60
	if minutes < 5 {
		return fmt.Sprintf("this should take 2-3 minutes")
	}

	if minutes < 10 {
		return fmt.Sprintf("this should take 5-10 minutes")
	}

	return fmt.Sprintf("this may take %d+ minutes", minutes)
}
