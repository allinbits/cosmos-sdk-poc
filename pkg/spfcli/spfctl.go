package spfcli

import "github.com/spf13/cobra"

func NewCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "spfcli",
		Long: "cli to interact with starport framework chains",
	}
	cliCtx := NewCLIContext()
	cmd.AddCommand(GetCmd(cliCtx))
	cmd.AddCommand(ResourcesCmd(cliCtx))
	return cmd
}
