package spfcli

import (
	"github.com/fdymylja/tmos/runtime/meta"
	"github.com/spf13/cobra"
)

func GetCmd(cli CLIContext) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "fetch a state object in the apiserver",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			client := cli.Client()

			resp, err := client.Get(cli.Context(), meta.APIGroup(args[0]), meta.APIKind(args[1]), args[2])
			if err != nil {
				return err
			}
			cmd.Printf("%s\n", resp)
			return nil
		},
	}
}
