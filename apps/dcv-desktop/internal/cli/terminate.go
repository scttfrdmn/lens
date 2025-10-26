package cli

import ("github.com/spf13/cobra")

func NewTerminateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "terminate [instance-id]",
		Short: "Terminate a DCV Desktop instance",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop terminate coming soon!")
			return nil
		},
	}
}
