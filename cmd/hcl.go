package cmd

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
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
	// obj, err := myhcl.HclToJson(filepath)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(obj)

	config := `
        resource "aws_instance" "web" {
            ami           = "ami-0c55b159cbfafe1f0"
            instance_type = "t2.micro"
        }
    `
	file, err := hclsyntax.ParseConfig([]byte(config), "", hcl.Pos{Line: 1, Column: 1})
	if err != nil {
		panic(err)
	}
	var result interface{}
	gohcl.DecodeBody(file.Body, nil, &result)
	if err != nil {
		panic(err)
	}
	print(string(file.Bytes))

	// // 抽象構文木から JSON 文字列を生成
	// jsonBytes, err := file.Body
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(jsonBytes))
	return nil
}

func Dumps(filepath string) error {
	// obj, err := myhcl.JsonToHcl(filepath)
	// if err != nil {
	// 	return err
	// }

	// fmt.Println(obj)
	return nil
}
