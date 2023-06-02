package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mo3789530/tfupdate/pkg/hcl"
	"github.com/spf13/cobra"
)

func HclCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "hcl",
		Short: "hcl commmand",
		Run: func(cmd *cobra.Command, args []string) {
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			switch args[0] {
			case "JsonToHcl":
				return toJson(args[1])
			case "HclToJson":
				return toHcl(args[1])
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

	command.Flags().StringP("filepath", "f", "", "filepath")
	return command
}

func toJson(filepath string) error {
	data, err := fileOpen(filepath)
	if err != nil {
		return err
	}
	hcl.JsonToHcl(string(data))
	return nil
}

func toHcl(filepath string) error {
	data, err := fileOpen(filepath)
	if err != nil {
		return err
	}
	print(hcl.HclBytesToJson(data))
	return nil
}

func fileOpen(filepath string) ([]byte, error) {
	return os.ReadFile(filepath)
}
