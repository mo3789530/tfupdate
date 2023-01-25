package hcl

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

// const (
// 	filename = "terraform.tf"
// )

type TerraformVersion struct {
}

func GetVersions(f *hclwrite.File) (string, error) {

	for _, tf := range findMatchingBlocks(f.Body(), "terraform", []string{}) {
		if tf.Body().GetAttribute("required_version") != nil {
		}
	}

	return "", nil
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
