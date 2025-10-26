package cli

import ("github.com/spf13/cobra")

func NewLaunchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "launch",
		Short: "Launch a new DCV Desktop instance",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop launch coming soon!")
			return nil
		},
	}
}
