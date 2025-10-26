package cli

import ("github.com/spf13/cobra")

func NewConnectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "connect [instance-id]",
		Short: "Connect to a running QGIS instance",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("QGIS connect coming soon!")
			return nil
		},
	}
}
