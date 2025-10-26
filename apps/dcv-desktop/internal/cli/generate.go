package cli

import ("github.com/spf13/cobra")

func NewGenerateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "generate",
		Short: "Generate cloud-init configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop generate coming soon!")
			return nil
		},
	}
}
