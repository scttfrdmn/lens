package cli

import (
	"github.com/spf13/cobra"
)

// NewWizardCmd creates the interactive wizard command for launching DCV Desktop
func NewWizardCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wizard",
		Short: "Interactive wizard to launch DCV Desktop environment",
		Long: `Interactive setup wizard that guides you through launching a NICE DCV Desktop instance.

The wizard will ask you about:
- Desktop environment (general-desktop, gpu-workstation, etc.)
- Instance type and size
- GPU requirements
- Storage size
- Auto-stop settings

This is the recommended way to launch your first DCV Desktop instance.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO: Implement wizard workflow
			cmd.Println("DCV Desktop wizard coming soon!")
			cmd.Println("This will guide you through launching a full Linux desktop with NICE DCV.")
			return nil
		},
	}

	return cmd
}
