package cli

import ("github.com/spf13/cobra")

func NewConnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "connect [instance-id]",
		Short: "Connect to a running DCV Desktop instance",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop connect coming soon!")
			return nil
		},
	}
}
