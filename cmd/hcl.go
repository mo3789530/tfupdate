package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func NewHCLCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "hcl",
		Short: "Run hcl command",
		// Run: func(cmd *cobra.Command, args []string) {
		// },
		RunE: func(cmd *cobra.Command, args []string) error {
			// var err error
			// dirs, _ := cmd.Flags().GetStringSlice("dirs")
			// relative, _ := cmd.Flags().GetString("relative")

			switch args[0] {
			case "json":
			case "hcl":
			default:
				log.Printf("%s not found \n", args[0])
			}

			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("<path> argument required")
			}
			return nil
		},
	}
	// command.Flags().StringSliceP("dirs", "d", []string{""}, "working dir")
	// command.Flags().StringP("relative", "r", "", "relative")

	return command
}
