package output

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Spinner displays an animated spinner for long-running operations
type Spinner struct {
	writer      io.Writer
	description string
	frames      []string
	interval    time.Duration
	active      bool
	mu          sync.Mutex
	stopCh      chan struct{}
	noColor     bool
}

// NewSpinner creates a new spinner with default frames
func NewSpinner(description string) *Spinner {
	return &Spinner{
		writer:      os.Stdout,
		description: description,
		frames:      []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		interval:    100 * time.Millisecond,
		stopCh:      make(chan struct{}),
		noColor:     false,
	}
}

// NewSimpleSpinner creates a spinner with ASCII-only frames
func NewSimpleSpinner(description string) *Spinner {
	return &Spinner{
		writer:      os.Stdout,
		description: description,
		frames:      []string{"|", "/", "-", "\\"},
		interval:    100 * time.Millisecond,
		stopCh:      make(chan struct{}),
		noColor:     false,
	}
}

// SetNoColor disables color output
func (s *Spinner) SetNoColor(noColor bool) {
	s.noColor = noColor
}

// SetFrames sets custom spinner frames
func (s *Spinner) SetFrames(frames []string) {
	s.frames = frames
}

// Start starts the spinner animation
func (s *Spinner) Start() {
	s.mu.Lock()
	if s.active {
		s.mu.Unlock()
		return
	}
	s.active = true
	s.mu.Unlock()

	go s.animate()
}

// animate runs the spinner animation loop
func (s *Spinner) animate() {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	frameIndex := 0

	for {
		select {
		case <-s.stopCh:
			// Clear the spinner line
			fmt.Fprintf(s.writer, "\r%s\r", clearLine())
			return
		case <-ticker.C:
			s.mu.Lock()
			if !s.active {
				s.mu.Unlock()
				return
			}

			frame := s.frames[frameIndex%len(s.frames)]
			var output string
			if s.noColor {
				output = fmt.Sprintf("\r%s %s", frame, s.description)
			} else {
				output = fmt.Sprintf("\r%s%s%s %s", colorCyan, frame, colorReset, s.description)
			}

			fmt.Fprint(s.writer, output)
			frameIndex++
			s.mu.Unlock()
		}
	}
}

// UpdateMessage updates the spinner's description while it's running
func (s *Spinner) UpdateMessage(description string) {
	s.mu.Lock()
	s.description = description
	s.mu.Unlock()
}

// Stop stops the spinner animation
func (s *Spinner) Stop() {
	s.mu.Lock()
	if !s.active {
		s.mu.Unlock()
		return
	}
	s.active = false
	s.mu.Unlock()

	close(s.stopCh)
	s.stopCh = make(chan struct{}) // Reset for potential reuse
}

// StopWithMessage stops the spinner and prints a final message
func (s *Spinner) StopWithMessage(message string) {
	s.Stop()
	fmt.Fprintln(s.writer, message)
}

// StopWithSuccess stops the spinner and prints a success message
func (s *Spinner) StopWithSuccess(message string) {
	s.Stop()
	formatter := NewFormatter(s.writer, s.noColor)
	formatter.Success(message)
}

// StopWithError stops the spinner and prints an error message
func (s *Spinner) StopWithError(message string) {
	s.Stop()
	formatter := NewFormatter(s.writer, s.noColor)
	formatter.Error(message)
}

// clearLine returns a string that clears the current terminal line
func clearLine() string {
	return "\033[2K"
}
