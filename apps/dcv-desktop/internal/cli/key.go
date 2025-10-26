package cli

import ("github.com/spf13/cobra")

func NewKeyCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "key",
		Short: "Manage SSH key pairs",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop key management coming soon!")
			return nil
		},
	}
}
