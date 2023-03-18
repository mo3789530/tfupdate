/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tfupdate/pkg/terraform"

	"github.com/spf13/cobra"
)

func NewCreateDocCommand() *cobra.Command {
	var createDocCmd = &cobra.Command{
		Use:   "createDoc",
		Short: "Create doc command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("createDoc called")
			t, _ := cmd.Flags().GetString("type")
			path, _ := cmd.Flags().GetString("path")
			// jsonData, _ := cmd.Flags().GetString("jsonDaa")
			switch t {
			case "plan":
				state := terraform.NewTfState("md")
				state.CreateDoc(path, nil)
			case "state":

			}
		},
	}

	createDocCmd.Flags().StringP("type", "t", "", "plan or state")
	createDocCmd.Flags().StringP("path", "p", "", "file path")
	return createDocCmd

}
