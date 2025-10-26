package cli

import ("github.com/spf13/cobra")

func NewDeleteAMICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete-ami",
		Short: "Delete a custom AMI",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop delete-ami coming soon!")
			return nil
		},
	}
}
