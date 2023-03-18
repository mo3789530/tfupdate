package hcl

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

type Config struct {
	Meta     Meta      `hcl:"terraform,block"`
	Provider *Provider `hcl:"provider,block"`
}

type Provider struct {
	Name string `hcl:"name,label"`
	Host string `hcl:"host,optional"`
}

type Meta struct {
	TfVersion    string       `hcl:"required_version"`
	ReqProviders *ReqProvider `hcl:"required_providers,block"`
}

type ReqProvider struct {
	Libvirt ProviderInfo `hcl:"docker,optional"`
	AWS     ProviderInfo `hcl:"aws,optional"`
}

type ProviderInfo struct {
	Source  string `cty:"source"`
	Version string `cty:"version"`
}

// https://github.com/orgrim/carcass/blob/fdbbcb85bf4e7628c7dd2c6f5a2b3814726722a6/terraform/config.go

func GetVersions(filepath string) (string, error) {
	src, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("error read file: %s", err)
		return "", err
	}

	file, diags := hclwrite.ParseConfig(src, filepath, hcl.InitialPos)
	if diags.HasErrors() {
		log.Printf("error hcl parse : %s", diags.Errs()[0])
		return "", diags.Errs()[0]
	}

	for _, tf := range findMatchingBlocks(file.Body(), "terraform", []string{}) {
		if tf.Body().GetAttribute("required_version") != nil {
			v := tf.Body().GetAttribute("required_version").Expr().BuildTokens(nil).Bytes()
			return string(v), nil
		}
	}
	return "", fmt.Errorf("error not found terrform block")
}

func findMatchingBlocks(b *hclwrite.Body, name string, labels []string) []*hclwrite.Block {

	var matched []*hclwrite.Block
	for _, block := range b.Blocks() {
		if name == block.Type() {
			labelsName := block.Labels()
			if len(labels) == 0 && len(labelsName) == 0 {
				matched = append(matched, block)
				continue
			}
			if reflect.DeepEqual(labels, name) {
				matched = append(matched, block)
			}
		}
	}
	return matched
}
