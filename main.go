package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	myhcl "tfupdate/pkg/hcl"

	"github.com/hashicorp/hcl/v2"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

var version string
var dir string

type Config struct {
	Meta Meta `hcl:"terraform,block"`
}

type Meta struct {
	TfVersion string `hcl:"required_version"`
	// ReqProviders ReqProvider `hcl:"required_providers,block"`
}

func main() {

	flag.StringVar(&version, "version", "0.14.0", "version")
	flag.StringVar(&dir, "dir", "./test/nullresource/", "dir")
	flag.Parse()
	println(version, dir)

	filepath := dir + "terraform.tf"
	src, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("error: %s", err)
	}

	var config Config

	// parser := hclparse.NewParser()
	// f, diags := parser.ParseHCLFile(filepath)
	// log.Print(diags)
	// moreDiafs := gohcl.DecodeBody(f.Body, nil, &config)
	// log.Print(moreDiafs)

	// https://github.com/orgrim/carcass/blob/fdbbcb85bf4e7628c7dd2c6f5a2b3814726722a6/terraform/config.go

	// err = hclsimple.DecodeFile(filepath, nil, &config)
	// if err != nil {
	// 	log.Fatalf("Failed to load configuration: %s", err)
	// }
	// log.Printf("Configuration is %#v", config)

	file, diags := hclwrite.ParseConfig(src, filepath, hcl.InitialPos)
	if diags.HasErrors() {
		log.Printf("err: %s", err)
	}
	log.Print(fmt.Sprintf("%s", file.Bytes()))
	hclsimple.Decode("dummy.hcl", file.Bytes(), nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %s", config.Meta.TfVersion)

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
