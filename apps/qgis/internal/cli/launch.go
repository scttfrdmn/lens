package cli

import ("github.com/spf13/cobra")

func NewLaunchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "launch",
		Short: "Launch a new QGIS instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("QGIS launch coming soon!")
			return nil
		},
	}
}
