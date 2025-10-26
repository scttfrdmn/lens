package cli

import ("github.com/spf13/cobra")

func NewCostsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "costs",
		Short: "Show cost estimates and usage",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("QGIS costs coming soon!")
			return nil
		},
	}
}
