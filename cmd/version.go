package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.0.1-rc"

func NewVersionCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "version",
		Short: "Print tfupdate version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(cmd.OutOrStdout(), "tfupdate %s \n", version)
		},
	}

	return command
}
