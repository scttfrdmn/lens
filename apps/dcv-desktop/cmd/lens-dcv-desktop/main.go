package main

import (
	"fmt"
	"os"

	"github.com/scttfrdmn/lens/apps/dcv-desktop/internal/cli"
	"github.com/scttfrdmn/lens/pkg"
	"github.com/scttfrdmn/lens/pkg/config"
	"github.com/spf13/cobra"
)

var (
	version = "0.10.0" // App version
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
		Use:   "lens-dcv-desktop",
		Short: "NICE DCV Desktop on AWS with GPU support for GUI research applications",
		Long: `lens-dcv-desktop provides a full Linux desktop environment via NICE DCV for GUI-based
research applications that require visual interfaces, 3D rendering, or GPU acceleration.

Features:
• Browser-based remote desktop with NICE DCV (no client required)
• GPU acceleration support for visualization and computation
• Session Manager secure connections (no exposed ports)
• Multiple desktop environments for different research domains
• Auto-stop on idle for cost optimization
• Support for MATLAB, QGIS, ParaView, ImageJ, and other GUI tools

Quick Start:
• Run 'lens-dcv-desktop' to launch the interactive setup wizard
• Or use 'lens-dcv-desktop quickstart' for instant launch with defaults
• Run 'lens-dcv-desktop --help' to see all available commands

Desktop Environments:
• general-desktop - Ubuntu desktop with research tools
• gpu-workstation - CUDA, visualization tools (requires GPU instance)
• matlab-desktop - MATLAB with full GUI support
• data-viz-desktop - ParaView, visualization tools
• image-analysis - ImageJ, Fiji, QuPath
• bioinformatics-gui - Geneious, UGENE, bioinformatics tools`,
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
