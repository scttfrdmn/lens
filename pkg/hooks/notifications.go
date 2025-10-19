package hooks

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/scttfrdmn/aws-ide/pkg/config"
)

// EventType represents different events that can trigger hooks
type EventType string

const (
	EventLaunchStarted   EventType = "launch_started"
	EventLaunchCompleted EventType = "launch_completed"
	EventLaunchFailed    EventType = "launch_failed"
	EventStopStarted     EventType = "stop_started"
	EventStopCompleted   EventType = "stop_completed"
	EventStopFailed      EventType = "stop_failed"
	EventConnectStarted  EventType = "connect_started"
	EventConnectFailed   EventType = "connect_failed"
)

// EventData contains information about the event
type EventData struct {
	EventType    EventType
	InstanceID   string
	InstanceType string
	Environment  string
	Region       string
	Timestamp    time.Time
	Error        string // Only populated for failed events
	AppName      string // "jupyter", "rstudio", "vscode"
}

// ExecuteHook runs a configured notification hook if one exists
func ExecuteHook(event EventData) error {
	// Check if hooks are configured
	cfg, err := config.LoadUserConfig()
	if err != nil {
		// No config or error reading - skip hooks silently
		return nil
	}

	// Get the hook command for this event type
	hookCmd := getHookCommand(cfg.Hooks, event.EventType)
	if hookCmd == "" {
		// No hook configured for this event - that's fine
		return nil
	}

	// Prepare environment variables for the hook
	env := prepareHookEnvironment(event)

	// Execute the hook command with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sh", "-c", hookCmd)
	cmd.Env = append(os.Environ(), env...)

	// Capture output but don't block on it
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Hook failed - log but don't fail the main operation
		fmt.Fprintf(os.Stderr, "Note: Notification hook failed: %v\n", err)
		if len(output) > 0 {
			fmt.Fprintf(os.Stderr, "Hook output: %s\n", string(output))
		}
		return err
	}

	return nil
}

// getHookCommand returns the configured hook command for an event type
func getHookCommand(hooks *config.HooksConfig, eventType EventType) string {
	if hooks == nil {
		return ""
	}

	switch eventType {
	case EventLaunchStarted:
		return hooks.OnLaunchStarted
	case EventLaunchCompleted:
		return hooks.OnLaunchCompleted
	case EventLaunchFailed:
		return hooks.OnLaunchFailed
	case EventStopStarted:
		return hooks.OnStopStarted
	case EventStopCompleted:
		return hooks.OnStopCompleted
	case EventStopFailed:
		return hooks.OnStopFailed
	case EventConnectStarted:
		return hooks.OnConnectStarted
	case EventConnectFailed:
		return hooks.OnConnectFailed
	default:
		return ""
	}
}

// prepareHookEnvironment creates environment variables for the hook
func prepareHookEnvironment(event EventData) []string {
	env := []string{
		fmt.Sprintf("AWS_IDE_EVENT=%s", event.EventType),
		fmt.Sprintf("AWS_IDE_INSTANCE_ID=%s", event.InstanceID),
		fmt.Sprintf("AWS_IDE_INSTANCE_TYPE=%s", event.InstanceType),
		fmt.Sprintf("AWS_IDE_ENVIRONMENT=%s", event.Environment),
		fmt.Sprintf("AWS_IDE_REGION=%s", event.Region),
		fmt.Sprintf("AWS_IDE_TIMESTAMP=%s", event.Timestamp.Format(time.RFC3339)),
		fmt.Sprintf("AWS_IDE_APP=%s", event.AppName),
	}

	if event.Error != "" {
		env = append(env, fmt.Sprintf("AWS_IDE_ERROR=%s", event.Error))
	}

	return env
}

// TriggerAsync executes a hook asynchronously (fire and forget)
func TriggerAsync(event EventData) {
	go func() {
		// Silently ignore errors in async mode
		_ = ExecuteHook(event)
	}()
}

// Example hook commands that users can configure:
//
// # Send email via sendmail
// echo "Instance ${AWS_IDE_INSTANCE_ID} started" | mail -s "AWS IDE Alert" user@example.com
//
// # Post to Slack webhook
// curl -X POST -H 'Content-type: application/json' \
//   --data "{\"text\":\"Instance ${AWS_IDE_INSTANCE_ID} ${AWS_IDE_EVENT}\"}" \
//   https://hooks.slack.com/services/YOUR/WEBHOOK/URL
//
// # Log to file
// echo "[$(date)] ${AWS_IDE_EVENT}: ${AWS_IDE_INSTANCE_ID}" >> ~/aws-ide-events.log
//
// # Send desktop notification (macOS)
// osascript -e "display notification \"${AWS_IDE_EVENT}\" with title \"AWS IDE\""
//
// # Send desktop notification (Linux)
// notify-send "AWS IDE" "${AWS_IDE_EVENT}: ${AWS_IDE_INSTANCE_ID}"

// FormatEventMessage creates a human-readable message for the event
func FormatEventMessage(event EventData) string {
	var action string
	switch event.EventType {
	case EventLaunchStarted:
		action = "is launching"
	case EventLaunchCompleted:
		action = "is ready"
	case EventLaunchFailed:
		action = "failed to launch"
	case EventStopStarted:
		action = "is stopping"
	case EventStopCompleted:
		action = "has stopped"
	case EventStopFailed:
		action = "failed to stop"
	case EventConnectStarted:
		action = "connection established"
	case EventConnectFailed:
		action = "connection failed"
	default:
		action = string(event.EventType)
	}

	msg := fmt.Sprintf("Instance %s %s", event.InstanceID, action)
	if event.Error != "" {
		msg += fmt.Sprintf(": %s", event.Error)
	}
	return msg
}
