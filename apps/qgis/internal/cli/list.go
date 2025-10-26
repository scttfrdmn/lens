package cli

import ("github.com/spf13/cobra")

func NewListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all QGIS instances",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("QGIS list coming soon!")
			return nil
		},
	}
}
