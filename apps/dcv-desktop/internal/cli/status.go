package cli

import ("github.com/spf13/cobra")

func NewStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status [instance-id]",
		Short: "Show status of DCV Desktop instance",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop status coming soon!")
			return nil
		},
	}
}
