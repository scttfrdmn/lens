package cli

import ("github.com/spf13/cobra")

func NewCreateAMICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create-ami",
		Short: "Create a custom AMI from running instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop create-ami coming soon!")
			return nil
		},
	}
}
