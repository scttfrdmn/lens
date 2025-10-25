package main

import (
	"fmt"
	"os"

	"github.com/scttfrdmn/lens/apps/vscode/internal/cli"
	"github.com/scttfrdmn/lens/pkg"
	"github.com/scttfrdmn/lens/pkg/config"
	"github.com/spf13/cobra"
)

var (
	version = "0.7.2" // App version
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
		Use:   "lens-vscode",
		Short: "VSCode Server on AWS Graviton with Session Manager & advanced networking",
		Long: `lens-vscode is a powerful CLI tool for launching VSCode Server (code-server) instances
on AWS EC2 Graviton processors with professional-grade security and networking.

Features:
• Full VSCode experience in your browser
• Session Manager & SSH connection methods
• Private subnet support with NAT Gateway
• Smart security groups and key management
• Built-in environments for web, Python, Go, and fullstack development
• Cost-aware infrastructure with reuse strategies
• Automatic extension installation

Quick Start:
• Just run 'lens-vscode' to launch the interactive setup wizard
• Or use 'lens-vscode quickstart' for instant launch with defaults
• Run 'lens-vscode --help' to see all available commands`,
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
	rootCmd.AddCommand(cli.NewStatusCmd())
	rootCmd.AddCommand(cli.NewEnvCmd())
	rootCmd.AddCommand(cli.NewKeyCmd())
	rootCmd.AddCommand(cli.NewCreateAMICmd())
	rootCmd.AddCommand(cli.NewDeleteAMICmd())
	rootCmd.AddCommand(cli.NewListAMIsCmd())
	rootCmd.AddCommand(cli.NewGenerateCmd())
	rootCmd.AddCommand(cli.NewConfigCmd())
	rootCmd.AddCommand(cli.NewCostsCmd())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
