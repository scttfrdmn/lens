package cli

import ("github.com/spf13/cobra")

func NewListAMIsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list-amis",
		Short: "List available DCV Desktop AMIs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop list-amis coming soon!")
			return nil
		},
	}
}
