package cli

import ("github.com/spf13/cobra")

func NewStartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start [instance-id]",
		Short: "Start a stopped DCV Desktop instance",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop start coming soon!")
			return nil
		},
	}
}
