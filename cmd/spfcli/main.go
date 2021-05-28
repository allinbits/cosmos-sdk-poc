package main

import (
	"fmt"
	"os"

	"github.com/fdymylja/tmos/pkg/spfcli"
)

func main() {
	cmd := spfcli.NewCLI()
	err := cmd.Execute()
	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
