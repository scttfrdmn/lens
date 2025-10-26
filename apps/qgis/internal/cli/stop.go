package cli

import ("github.com/spf13/cobra")

func NewStopCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stop [instance-id]",
		Short: "Stop a running QGIS instance",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("QGIS stop coming soon!")
			return nil
		},
	}
}
