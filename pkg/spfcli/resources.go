package spfcli

import (
	"github.com/spf13/cobra"
)

func ResourcesCmd(cli CLIContext) *cobra.Command {
	return &cobra.Command{
		Use:   "resources",
		Short: "lists the chain's resources",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := cli.Client()

			resp, err := client.Resources(cli.Context())
			if err != nil {
				return err
			}
			cmd.Printf("%s\n", resp)
			return nil
		},
	}
}
