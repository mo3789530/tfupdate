package main

import (
	"flag"
	"log"
	"strings"
	myhcl "tfupdate/pkg/hcl"
	"tfupdate/pkg/terraform"
)

var version string
var dir string
var relative string

func main() {

	flag.StringVar(&version, "version", "0.14.0", "version")
	flag.StringVar(&dir, "dir", "mongo", "dir")
	flag.StringVar(&relative, "relative", "test", "relative")
	flag.Parse()
	// println(version, dir)

	dirs := strings.Split(dir, " ")
	for _, v := range dirs {
		folderpath := relative + "/" + v
		filepath := folderpath + "/terraform.tf"
		log.Print(filepath)
		version, err := myhcl.GetVersions(filepath)
		if err != nil {
			log.Fatalf("err %s", err)
		}
		log.Printf("%s using terrafrom version: %s", v, version)
		exec := terraform.NewExec(version, true)
		tf, err := exec.Init(v)
		if err != nil {
			log.Printf("error init %s", err)
		}
		isdiffer, err := exec.Plan(tf)
		if err != nil {
			log.Printf("err plan %s", err)
		}
		if !isdiffer {
			log.Printf("no changes")
			continue
		}
		exec.Show(tf, false)
	}

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
