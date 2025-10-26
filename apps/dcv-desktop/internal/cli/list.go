package cli

import ("github.com/spf13/cobra")

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all DCV Desktop instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop list coming soon!")
			return nil
		},
	}
}
