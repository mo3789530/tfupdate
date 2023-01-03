package main

import (
	"flag"
	"log"
	"tfupdate/pkg/terraform"
	"tfupdate/pkg/utils"
)

var version string
var dir string

func main() {

	flag.StringVar(&version, "version", "", "version")
	flag.StringVar(&dir, "dir", "./", "dir")
	flag.Parse()
	println(version, dir)

	futils := utils.NewFolderUtils(dir)
	folders := futils.ListDir()

	for _, f := range folders {
		tfe := terraform.NewExec(version)
		tf, err := tfe.Init(dir + "/" + f)
		if err != nil {
			log.Fatalf("error: %s", err)
		}
		tfe.Plan(tf)
	}
}
