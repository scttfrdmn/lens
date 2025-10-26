package cli

import ("github.com/spf13/cobra")

func NewEnvCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "env",
		Short: "List available desktop environments",
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.Println("DCV Desktop environments coming soon!")
			return nil
		},
	}
}
