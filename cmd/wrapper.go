package cmd

import (
	"fmt"
	"log"
	"regexp"
	"tfupdate/pkg/terraform"

	myhcl "tfupdate/pkg/hcl"

	"github.com/spf13/cobra"
)

func NewWrapperCommand() *cobra.Command {
	var command = &cobra.Command{
		Use:   "tf",
		Short: "Run terraform command",
		// Run: func(cmd *cobra.Command, args []string) {
		// },
		RunE: func(cmd *cobra.Command, args []string) error {
			// var err error
			fmt.Println(args)
			dirs, _ := cmd.Flags().GetStringSlice("dirs")
			relative, _ := cmd.Flags().GetString("relative")

			switch args[0] {
			case "plan":
				//fmt.Println(dirs)
				//fmt.Println(relative)
				return runPlan(dirs, relative)
			case "apply":
				runApply(dirs, relative)
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
	command.Flags().StringSliceP("dirs", "d", []string{""}, "working dir")
	command.Flags().StringP("relative", "r", "", "relative")

	return command
}

func runPlan(dirs []string, relative string) error {
	for _, v := range dirs {
		folderpath := relative + v
		filepath := folderpath + "/terraform.tf"

		log.Printf("target path: %s \n", filepath)

		strVersion, err := myhcl.GetVersions(filepath)
		if err != nil {
			log.Fatalf("err %s", err)
			return err
		}
		rex := regexp.MustCompile("[0-9.]+")
		version = rex.FindString(strVersion)

		log.Printf("%s using terrafrom version: %s", v, version)

		exec := terraform.NewExec(version, true)
		tf, err := exec.Init(folderpath)
		if err != nil {
			log.Printf("error init %s", err)
			return err
		}
		isdiffer, err := exec.Plan(tf)
		if err != nil {
			log.Printf("err plan %s", err)
			return err
		}
		if !isdiffer {
			log.Printf("no changes")
			continue
		} else {
			show, err := exec.Show(tf, false)
			if err != nil {
				log.Printf("err show %s", err)
				return err
			}
			log.Println(show)
		}
	}
	return nil
}

func runApply(dirs []string, relative string) error {
	return nil
}
