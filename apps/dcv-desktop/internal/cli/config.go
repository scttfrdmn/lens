package cli

import ("github.com/spf13/cobra")

func NewConfigCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Manage lens configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop config coming soon!")
			return nil
		},
	}
}
