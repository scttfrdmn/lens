package cli

import ("github.com/spf13/cobra")

func NewListAMIsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list-amis",
		Short: "List available QGIS AMIs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("QGIS list-amis coming soon!")
			return nil
		},
	}
}
