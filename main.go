package main

import (
	"flag"
	"io/ioutil"
	"log"
	myhcl "tfupdate/pkg/hcl"

	"github.com/hashicorp/hcl/v2"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

var version string
var dir string

func main() {

	flag.StringVar(&version, "version", "0.14.0", "version")
	flag.StringVar(&dir, "dir", "./test/nullresource/", "dir")
	flag.Parse()
	println(version, dir)

	filepath := dir + "/terraform.tf"
	src, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Println("error: %s", err)
	}

	file, diags := hclwrite.ParseConfig(src, filepath, hcl.InitialPos)
	if diags.HasErrors() {
		log.Printf("err: %s", err)
	}

	log.Print(file)

	myhcl.GetVersions(file)

	// futils := utils.NewFolderUtils(dir)
	// folders := futils.ListDir()

	// for _, f := range folders {
	// 	tfe := terraform.NewExec(version)
	// 	tf, err := tfe.Init(dir + "/" + f)
	// 	if err != nil {
	// 		log.Fatalf("error: %s", err)
	// 	}
	// 	tfe.Plan(tf)
	// }
	// tfe := terraform.NewExec(version, true)
	// tf, err := tfe.Init(dir)
	// if err != nil {
	// 	log.Fatalf("error: %s", err)
	// }
	// tfe.Plan(tf)
	// show, err := tfe.Show(tf, false)
	// if err != nil {
	// 	log.Fatalf("error: %s", err)
	// }
	// log.Println(show)

}
