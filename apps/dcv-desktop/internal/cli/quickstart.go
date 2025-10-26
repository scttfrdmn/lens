package cli

import (
	"github.com/spf13/cobra"
)

func NewQuickstartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "quickstart",
		Short: "Quick launch with sensible defaults",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop quickstart coming soon!")
			return nil
		},
	}
}
