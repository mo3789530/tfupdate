package cmd

import (
	"fmt"
	myhcl "tfupdate/pkg/hcl"

	"github.com/spf13/cobra"
)

func NewHclCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "hcl",
		Short: "Run terraform command",
		// Run: func(cmd *cobra.Command, args []string) {
		// },
		RunE: func(cmd *cobra.Command, args []string) error {
			// var err error
			path, _ := cmd.Flags().GetString("path")

			switch args[0] {
			case "loads":
				//fmt.Println(dirs)
				//fmt.Println(relative)
				return Loads(path)
			case "dumps":
				return Dumps(path)
			default:
				return fmt.Errorf("%s not found", args[0])
			}

		},
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("<path> argument required")
			}
			return nil
		},
	}
	command.Flags().StringP("path", "p", "", "file path")
	// command.Flags().StringP("relative", "r", "", "relative")

	return command
}

func Loads(filepath string) error {
	obj, err := myhcl.HclToJsonString(filepath)
	if err != nil {
		return err
	}

	fmt.Println(obj)
	return nil
}

func Dumps(filepath string) error {
	obj, err := myhcl.JsonToHcl(filepath)
	if err != nil {
		return err
	}

	fmt.Println(obj)
	return nil
}
