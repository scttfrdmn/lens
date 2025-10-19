package cli

import (
	"fmt"
	"time"

	"github.com/scttfrdmn/aws-ide/pkg/output"
	"github.com/spf13/cobra"
)

// NewQuickstartCmd creates the quickstart command for aws-vscode
func NewQuickstartCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quickstart",
		Short: "Launch VSCode Server with smart defaults (fastest way to get started)",
		Long: `The quickstart command launches a VSCode Server instance with sensible defaults,
perfect for getting started immediately without making any decisions.

This is the fastest way to launch - just one command and you're running!

Smart defaults used:
- Environment: default (Node.js, Python, Go development tools)
- Computer power: t4g.medium (balanced performance, ~$0.03/hour)
- Storage: 50GB (enough for most projects)
- Auto-stop: 2 hours of inactivity (saves money automatically)
- Name: Auto-generated based on date/time

You can always customize these later with the 'wizard' or 'launch' commands.`,
		Example: `  # Launch with smart defaults
  aws-vscode quickstart

  # Launch and connect immediately
  aws-vscode quickstart && aws-vscode connect

  # Launch in a specific region
  aws-vscode quickstart --region us-west-2`,
		RunE: runQuickstart,
	}

	// Allow optional overrides for advanced users
	cmd.Flags().String("region", "", "AWS region (default: your configured region)")
	cmd.Flags().Bool("dry-run", false, "Show what would be created without actually creating it")

	return cmd
}

func runQuickstart(cmd *cobra.Command, args []string) error {
	out := output.DefaultFormatter()

	out.Header("🚀 VSCode Server Quickstart")
	out.Blank()
	out.Info("Launching with smart defaults - no questions, no hassle!")
	out.Blank()

	// Get optional overrides
	region, _ := cmd.Flags().GetString("region")
	dryRun, _ := cmd.Flags().GetBool("dry-run")

	// Define quickstart defaults
	defaults := map[string]string{
		"environment":    "default",
		"instance-type":  "t4g.medium",
		"ebs-size":       "50",
		"idle-timeout":   "2h",
		"name":           generateQuickstartName(),
	}

	// Show what we're launching
	out.Subheader("Your Quickstart Configuration:")
	out.KeyValue("Environment", "Default (Node.js, Python, Go + Amazon Q Developer)")
	out.KeyValue("Computer Power", "t4g.medium (2 vCPU, 4GB RAM)")
	out.KeyValue("Storage", "50 GB SSD")
	out.KeyValue("Auto-Stop", "After 2 hours of inactivity")
	out.KeyValue("Instance Name", defaults["name"])
	if region != "" {
		out.KeyValue("Region", region)
	}
	out.Blank()

	// Show cost estimate
	out.Subheader("💰 Cost Estimate:")
	out.KeyValue("Hourly Cost", "$0.0336/hour")
	out.KeyValue("Daily Cost (8 hours)", "~$0.27/day")
	out.KeyValue("Monthly Cost (40 hours/week)", "~$5.38/month")
	out.Blank()
	out.Info("💡 Auto-stop will shut down after 2 hours idle to minimize costs")
	out.Blank()

	if dryRun {
		out.DryRun("This is a dry run - no instance will be created")
		out.Info("To actually launch, run without --dry-run flag")
		return nil
	}

	out.Progress("Launching your VSCode Server environment...")
	out.Blank()

	// Build the launch command with defaults
	launchCmd := NewLaunchCmd()

	// Set all the default flags
	launchCmd.Flags().Set("env", defaults["environment"])
	launchCmd.Flags().Set("instance-type", defaults["instance-type"])
	launchCmd.Flags().Set("ebs-size", defaults["ebs-size"])
	launchCmd.Flags().Set("idle-timeout", defaults["idle-timeout"])
	launchCmd.Flags().Set("name", defaults["name"])

	if region != "" {
		launchCmd.Flags().Set("region", region)
	}

	// Execute the launch command
	err := launchCmd.Execute()
	if err != nil {
		out.Error("Failed to launch quickstart instance")
		return err
	}

	out.Blank()
	out.Separator()
	out.Success("🎉 Quickstart launch complete!")
	out.Blank()
	out.Info("💡 Next steps:")
	out.List("Wait 2-3 minutes for VSCode to finish installing")
	out.List("Connect with: aws-vscode connect " + defaults["name"])
	out.List("When done, stop it to save money: aws-vscode stop " + defaults["name"])
	out.List("Or just leave it - auto-stop will shut it down after 2 hours")
	out.Blank()

	return nil
}

// generateQuickstartName creates a unique name for quickstart instances
func generateQuickstartName() string {
	now := time.Now()
	return fmt.Sprintf("vscode-quickstart-%s", now.Format("2006-01-02-150405"))
}
