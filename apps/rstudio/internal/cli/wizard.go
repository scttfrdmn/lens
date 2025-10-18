package cli

import (
	"fmt"

	"github.com/scttfrdmn/aws-ide/pkg/cli"
	"github.com/spf13/cobra"
)

// NewWizardCmd creates the wizard command for aws-rstudio
func NewWizardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wizard",
		Short: "Interactive setup wizard for launching RStudio Server",
		Long: `The wizard command provides an interactive, beginner-friendly way to launch
your first RStudio Server instance. Instead of remembering flags and options,
the wizard will ask you simple questions in plain language and set everything
up for you automatically.

Perfect for researchers who want to get started quickly without technical details.`,
		Example: `  # Launch the wizard
  aws-rstudio wizard

  The wizard will guide you through:
  - Choosing what type of analysis you want to do
  - Selecting the right computer power for your needs
  - Setting up storage space
  - Configuring auto-stop to save money
  - Giving your instance a name`,
		RunE: runWizard,
	}

	return cmd
}

func runWizard(cmd *cobra.Command, args []string) error {
	// Configure the wizard for RStudio
	config := cli.WizardConfig{
		AppName: "AWS RStudio",
		AppType: "rstudio",
	}

	// Run the interactive wizard
	result, err := cli.RunLaunchWizard(config)
	if err != nil {
		return fmt.Errorf("wizard cancelled: %w", err)
	}

	fmt.Println()
	fmt.Println("🚀 Launching your RStudio Server instance...")
	fmt.Println()

	// Build the launch command arguments from wizard results
	launchCmd := NewLaunchCmd()

	// Set flags based on wizard results
	launchCmd.Flags().Set("env", result.Environment)
	launchCmd.Flags().Set("instance-type", result.InstanceType)
	launchCmd.Flags().Set("ebs-size", fmt.Sprintf("%d", result.EBSSize))

	if result.IdleTimeout != "" {
		launchCmd.Flags().Set("idle-timeout", result.IdleTimeout)
	}

	if result.Name != "" {
		launchCmd.Flags().Set("name", result.Name)
	}

	// Execute the launch command
	return launchCmd.Execute()
}
