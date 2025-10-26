package cli

import (
	"github.com/spf13/cobra"
)

// NewWizardCmd creates the interactive wizard command for launching QGIS
func NewWizardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wizard",
		Short: "Interactive wizard to launch QGIS environment",
		Long: `Interactive setup wizard that guides you through launching a NICE QGIS instance.

The wizard will ask you about:
- Desktop environment (general-desktop, gpu-workstation, etc.)
- Instance type and size
- GPU requirements
- Storage size
- Auto-stop settings

This is the recommended way to launch your first QGIS instance.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Implement wizard workflow
			cmd.Println("QGIS wizard coming soon!")
			cmd.Println("This will guide you through launching a full Linux desktop with NICE DCV.")
			return nil
		},
	}

	return cmd
}
