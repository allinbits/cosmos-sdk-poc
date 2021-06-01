package spfcli

import (
	"github.com/spf13/cobra"
)

func GetCmd(cli CLIContext) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "fetch a state object in the apiserver",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := cli.Client()

			resp, err := client.Get(cli.Context(), args[0], args[1], args[2])
			if err != nil {
				return err
			}
			cmd.Printf("%s\n", resp)
			return nil
		},
	}
}
