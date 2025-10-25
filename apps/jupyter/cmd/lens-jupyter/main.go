package main

import (
	"fmt"
	"os"

	"github.com/scttfrdmn/lens/apps/jupyter/internal/cli"
	"github.com/scttfrdmn/lens/pkg"
	"github.com/scttfrdmn/lens/pkg/config"
	"github.com/spf13/cobra"
)

var (
	version = "0.9.0" // App version
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	// Migrate from legacy config directories if needed
	if err := config.MigrateFromLegacy(); err != nil {
		fmt.Fprintf(os.Stderr, "Warning: Failed to migrate legacy config: %v\n", err)
		// Continue anyway - don't block if migration fails
	}

	if err := config.EnsureConfigDir(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create config directory: %v\n", err)
		os.Exit(1)
	}

	rootCmd := &cobra.Command{
		Use:   "lens-jupyter",
		Short: "Secure Jupyter Lab on AWS Graviton with Session Manager & advanced networking",
		Long: `lens-jupyter is a powerful CLI tool for launching secure Jupyter Lab instances
on AWS EC2 Graviton processors with professional-grade security and networking.

Features:
• Session Manager & SSH connection methods
• Private subnet support with NAT Gateway
• Smart security groups and key management
• Built-in environments for data science, ML, and research
• Cost-aware infrastructure with reuse strategies

Quick Start:
• Just run 'lens-jupyter' to launch the interactive setup wizard
• Or use 'lens-jupyter quickstart' for instant launch with defaults
• Run 'lens-jupyter --help' to see all available commands`,
		Version: fmt.Sprintf("v%s (platform: v%s, commit: %s, date: %s)", version, pkg.Version, commit, date),
		Run: func(cmd *cobra.Command, args []string) {
			// When no subcommand is provided, run the wizard by default
			wizardCmd := cli.NewWizardCmd()
			wizardCmd.SetArgs(args)
			if err := wizardCmd.Execute(); err != nil {
				os.Exit(1)
			}
		},
	}

	// Add subcommands
	rootCmd.AddCommand(cli.NewQuickstartCmd())
	rootCmd.AddCommand(cli.NewWizardCmd())
	rootCmd.AddCommand(cli.NewLaunchCmd())
	rootCmd.AddCommand(cli.NewListCmd())
	rootCmd.AddCommand(cli.NewConnectCmd())
	rootCmd.AddCommand(cli.NewStopCmd())
	rootCmd.AddCommand(cli.NewStartCmd())
	rootCmd.AddCommand(cli.NewTerminateCmd())
	rootCmd.AddCommand(cli.NewCreateAMICmd())
	rootCmd.AddCommand(cli.NewListAMIsCmd())
	rootCmd.AddCommand(cli.NewDeleteAMICmd())
	rootCmd.AddCommand(cli.NewEnvCmd())
	rootCmd.AddCommand(cli.NewStatusCmd())
	rootCmd.AddCommand(cli.NewGenerateCmd())
	rootCmd.AddCommand(cli.NewKeyCmd())
	rootCmd.AddCommand(cli.NewConfigCmd())
	rootCmd.AddCommand(cli.NewCostsCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
