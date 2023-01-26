package hcl

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
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
	src, err := ioutil.ReadFile(filepath)
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

func getHCLNativeAttribute(body *hclwrite.Body, name string) (*hcl.Attribute, error) {
	attr := body.GetAttribute(name)
	if attr == nil {
		return nil, nil
	}

	// build low-level byte sequences
	attrAsBytes := attr.Expr().BuildTokens(nil).Bytes()
	src := append([]byte(name+" = "), attrAsBytes...)

	// parse an expression as a hcl.File.
	// Note that an attribute may contains references, which are defined outside the file.
	// So we cannot simply use hclsyntax.ParseExpression or hclsyntax.ParseConfig here.
	// We need to use a loe-level parser not to resolve all references.
	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(src, "generated_by_getHCLNativeAttribute")
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse expression: %s", diags)
	}

	attrs, diags := file.Body.JustAttributes()
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to get attributes: %s", diags)
	}

	hclAttr, ok := attrs[name]
	if !ok {
		return nil, fmt.Errorf("attribute not found: %s", src)
	}

	return hclAttr, nil
}
