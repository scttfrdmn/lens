package errors

import (
	"fmt"
	"strings"
)

// FriendlyError wraps technical errors with beginner-friendly messages
// designed for non-technical academic researchers
type FriendlyError struct {
	Title       string   // Short, plain-English title (e.g., "Can't connect to AWS")
	Explanation string   // Non-technical explanation of what went wrong
	Cause       error    // The underlying technical error (hidden from users by default)
	NextSteps   []string // Simple, actionable steps to fix the problem
	ShowCause   bool     // Whether to show the technical error (for debugging)
}

// Error implements the error interface
func (e *FriendlyError) Error() string {
	var sb strings.Builder

	// Title with emoji for visual clarity
	sb.WriteString(fmt.Sprintf("❌ %s\n\n", e.Title))

	// Plain-English explanation
	if e.Explanation != "" {
		sb.WriteString(fmt.Sprintf("%s\n", e.Explanation))
	}

	// Next steps
	if len(e.NextSteps) > 0 {
		sb.WriteString("\n💡 What to do next:\n")
		for i, step := range e.NextSteps {
			sb.WriteString(fmt.Sprintf("  %d. %s\n", i+1, step))
		}
	}

	// Show technical details if requested (for debugging)
	if e.ShowCause && e.Cause != nil {
		sb.WriteString(fmt.Sprintf("\n🔍 Technical details (for support): %v\n", e.Cause))
	}

	return sb.String()
}

// Unwrap returns the underlying error
func (e *FriendlyError) Unwrap() error {
	return e.Cause
}

// NewFriendly creates a beginner-friendly error
func NewFriendly(title string, explanation string, cause error) *FriendlyError {
	return &FriendlyError{
		Title:       title,
		Explanation: explanation,
		Cause:       cause,
		NextSteps:   []string{},
		ShowCause:   false,
	}
}

// WithNextStep adds a next step to the error
func (e *FriendlyError) WithNextStep(step string) *FriendlyError {
	e.NextSteps = append(e.NextSteps, step)
	return e
}

// WithNextSteps adds multiple next steps
func (e *FriendlyError) WithNextSteps(steps ...string) *FriendlyError {
	e.NextSteps = append(e.NextSteps, steps...)
	return e
}

// ShowTechnicalDetails enables display of technical error details
func (e *FriendlyError) ShowTechnicalDetails() *FriendlyError {
	e.ShowCause = true
	return e
}

// Common friendly error patterns for non-technical users

// AWSCredentialsError creates a friendly error for AWS credential issues
func AWSCredentialsError(cause error) *FriendlyError {
	return NewFriendly(
		"Can't connect to AWS",
		"Your AWS credentials aren't set up correctly. This is like trying to log into a website without a password.",
		cause,
	).WithNextSteps(
		"Run this command to set up your credentials: aws configure",
		"You'll need your AWS Access Key ID and Secret Access Key (ask your IT department if you don't have these)",
		"If you're not sure what to enter, see the AWS setup guide in the documentation",
	)
}

// InstanceNotRunningError creates a friendly error when instance isn't running
func InstanceNotRunningError(instanceID string, state string) *FriendlyError {
	var explanation string
	if state == "stopped" {
		explanation = "Your cloud computer is turned off right now. You need to start it before you can connect to it."
	} else if state == "stopping" {
		explanation = "Your cloud computer is currently shutting down. Wait a minute for it to fully stop, then you can start it again."
	} else if state == "pending" {
		explanation = "Your cloud computer is still starting up. This usually takes 2-3 minutes."
	} else {
		explanation = fmt.Sprintf("Your cloud computer is in '%s' state and can't be connected to right now.", state)
	}

	return NewFriendly(
		"Instance is not running",
		explanation,
		fmt.Errorf("instance %s is in %s state", instanceID, state),
	).WithNextSteps(
		fmt.Sprintf("Start the instance: lens-jupyter start %s", instanceID),
		"Wait 2-3 minutes for it to fully start",
		fmt.Sprintf("Check if it's ready: lens-jupyter status %s", instanceID),
	)
}

// NoInstancesFoundError creates a friendly error when user has no instances
func NoInstancesFoundError() *FriendlyError {
	return NewFriendly(
		"You don't have any cloud computers yet",
		"You haven't created any cloud computers (instances) yet. You need to create one first before you can use it.",
		fmt.Errorf("no instances found"),
	).WithNextSteps(
		"Create your first instance with the easy setup wizard: lens-jupyter wizard",
		"Or create one with specific options: lens-jupyter launch",
		"Need help? Run: lens-jupyter wizard --help",
	)
}

// InstanceTerminatedError creates a friendly error when instance was terminated
func InstanceTerminatedError(instanceID string) *FriendlyError {
	return NewFriendly(
		"That instance was deleted",
		"The cloud computer you're trying to use has been deleted (terminated). Once an instance is terminated, it's gone permanently and can't be recovered.",
		fmt.Errorf("instance %s was terminated", instanceID),
	).WithNextSteps(
		"Create a new instance with: lens-jupyter wizard",
		"Check your other instances: lens-jupyter list",
		"If you need to recover data, check if you have backups or snapshots",
	)
}

// RegionMismatchError creates a friendly error for wrong AWS region
func RegionMismatchError(instanceRegion string, currentRegion string) *FriendlyError {
	return NewFriendly(
		"Wrong AWS region",
		fmt.Sprintf("Your cloud computer is in the '%s' region, but you're currently looking in the '%s' region. AWS regions are like different data centers around the world.", instanceRegion, currentRegion),
		fmt.Errorf("instance is in %s but you're using %s", instanceRegion, currentRegion),
	).WithNextSteps(
		fmt.Sprintf("Switch to the correct region: lens-jupyter --region %s list", instanceRegion),
		fmt.Sprintf("Or set it as default: lens-jupyter config set default_region %s", instanceRegion),
		"List all your instances in all regions: lens-jupyter list --all-regions",
	)
}

// OutOfMoneyError creates a friendly error when billing issues occur
func OutOfMoneyError(cause error) *FriendlyError {
	return NewFriendly(
		"Can't create instance - billing issue",
		"AWS can't create a new cloud computer for you right now. This usually happens when there's a payment issue with your AWS account.",
		cause,
	).WithNextSteps(
		"Check your AWS billing dashboard to see if there are payment issues",
		"Make sure you have a valid payment method on file",
		"Contact AWS support if you think this is a mistake",
		"If you're using a free tier account, check if you've exceeded the free usage limits",
	)
}

// TooManyInstancesError creates a friendly error for quota limits
func TooManyInstancesError(cause error) *FriendlyError {
	return NewFriendly(
		"Too many instances running",
		"You've reached the maximum number of cloud computers (instances) you're allowed to have running at once. AWS limits how many you can run to prevent accidents.",
		cause,
	).WithNextSteps(
		"Stop or terminate instances you're not using: lens-jupyter list",
		"Wait for stopped instances to fully shut down (takes 1-2 minutes)",
		"If you need more instances, request a limit increase in the AWS console",
	)
}

// NetworkTimeoutError creates a friendly error for connection timeouts
func NetworkTimeoutError(cause error) *FriendlyError {
	return NewFriendly(
		"Connection timed out",
		"The connection to your cloud computer took too long and gave up. This usually means the computer is slow to respond or there's a network problem.",
		cause,
	).WithNextSteps(
		"Check that your internet connection is working",
		"Verify the instance is running: lens-jupyter status INSTANCE_ID",
		"Try again in a minute - the instance might still be starting up",
		"If this keeps happening, the instance may have crashed. Try stopping and starting it again.",
	)
}

// EnvironmentNotFoundFriendly creates a friendly error for missing environments
func EnvironmentNotFoundFriendly(envName string) *FriendlyError {
	return NewFriendly(
		"Environment not found",
		fmt.Sprintf("The software setup (environment) called '%s' doesn't exist. Environments are pre-configured sets of software tools.", envName),
		fmt.Errorf("environment '%s' not found", envName),
	).WithNextSteps(
		"See what environments are available: lens-jupyter env list",
		"Use the wizard to pick the right one: lens-jupyter wizard",
		"Create a custom environment: lens-jupyter generate",
	)
}

// DiskFullError creates a friendly error when storage is full
func DiskFullError(instanceID string, cause error) *FriendlyError {
	return NewFriendly(
		"Your cloud computer is out of storage space",
		"The hard drive on your cloud computer is completely full. You need to delete some files or increase the storage size.",
		cause,
	).WithNextSteps(
		fmt.Sprintf("Connect and delete unnecessary files: lens-jupyter connect %s", instanceID),
		"Check what's using space with: df -h",
		"Create a new instance with more storage space",
		"Consider creating a snapshot before deleting important files",
	)
}

// PermissionDeniedFriendly creates a friendly error for permission issues
func PermissionDeniedFriendly(action string, cause error) *FriendlyError {
	return NewFriendly(
		"You don't have permission to do that",
		fmt.Sprintf("Your AWS account doesn't allow you to %s. This is a security setting that your AWS administrator controls.", action),
		cause,
	).WithNextSteps(
		"Contact your AWS account administrator or IT department",
		"Tell them you need permission to: "+action,
		"Ask them to give you the necessary IAM (permission) policies",
	)
}

// TranslateAWSError attempts to convert technical AWS errors into friendly messages
func TranslateAWSError(err error, context string) error {
	if err == nil {
		return nil
	}

	errMsg := strings.ToLower(err.Error())

	// Check for common AWS error patterns
	switch {
	case strings.Contains(errMsg, "no credentials") || strings.Contains(errMsg, "unable to locate credentials"):
		return AWSCredentialsError(err)

	case strings.Contains(errMsg, "authoriz") || strings.Contains(errMsg, "access denied"):
		return PermissionDeniedFriendly(context, err)

	case strings.Contains(errMsg, "invalidinstanceid"):
		return InstanceTerminatedError(extractInstanceID(errMsg))

	case strings.Contains(errMsg, "timeout") || strings.Contains(errMsg, "timed out"):
		return NetworkTimeoutError(err)

	case strings.Contains(errMsg, "billing") || strings.Contains(errMsg, "payment"):
		return OutOfMoneyError(err)

	case strings.Contains(errMsg, "instancelimitexceeded") || strings.Contains(errMsg, "quota"):
		return TooManyInstancesError(err)

	case strings.Contains(errMsg, "no space left") || strings.Contains(errMsg, "disk full"):
		return DiskFullError("unknown", err)

	default:
		// Return a generic friendly error
		return NewFriendly(
			"Something went wrong",
			"An unexpected problem occurred. This might be temporary, so it's worth trying again.",
			err,
		).WithNextSteps(
			"Try the command again in a minute",
			"Check if your internet connection is working",
			"Run 'lens-jupyter status' to check if everything is okay",
			"If this keeps happening, you can report it as a bug",
		).ShowTechnicalDetails()
	}
}

// extractInstanceID attempts to extract an instance ID from an error message
func extractInstanceID(errMsg string) string {
	// Look for instance ID pattern: i-xxxxxxxxx
	parts := strings.Split(errMsg, "i-")
	if len(parts) > 1 {
		id := "i-" + strings.Fields(parts[1])[0]
		return strings.TrimRight(id, "'\":,.")
	}
	return "unknown"
}
