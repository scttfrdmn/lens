package main

import (
	"fmt"
	"os"

	"github.com/scttfrdmn/lens/apps/qgis/internal/cli"
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
		Use:   "lens-qgis",
		Short: "QGIS on AWS with NICE DCV remote desktop",
		Long: `lens-qgis launches QGIS (Geographic Information System) on AWS EC2 with
browser-based remote desktop access via NICE DCV.

Features:
• Full QGIS desktop experience in your browser
• No local installation required
• GPU acceleration for large raster datasets (optional)
• AWS Session Manager secure connections (no exposed ports)
• Built-in environments for different GIS workflows
• Auto-stop on idle for cost optimization

QGIS Environments:
• basic-gis - QGIS with essential plugins (t3.xlarge, ~$0.17/hr)
• advanced-gis - QGIS + GRASS + SAGA + PostGIS (t3.xlarge)
• remote-sensing - QGIS + Orfeo Toolbox + SNAP + GPU (g4dn.xlarge, ~$0.53/hr)

Quick Start:
• Run 'lens-qgis' to launch the interactive setup wizard
• Or use 'lens-qgis quickstart' for instant launch with defaults
• Run 'lens-qgis --help' to see all available commands`,
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
