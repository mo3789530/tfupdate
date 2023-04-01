package hcl

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func HclToJson() {

}

func JsonToHcl() {

}

func FindMatchingBlocks(b *hclwrite.Body, name string, labels []string) []*hclwrite.Block {
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

// write block
func WriteBodyBodyHcl(body *hclwrite.Body, key string, values []interface{}) {
	if len(values) < 1 {
		return
	}

	block := body.AppendNewBlock(key, []string{})
	blockBody := block.Body()
	for _, value := range values {
		WriteAttribute(blockBody, value)
	}

}

// write attribute
func WriteAttribute(blockBody *hclwrite.Body, value interface{}) {
	switch vt := value.(type) {
	case map[string]interface{}:
		for k, v := range vt {
			switch vv := v.(type) {
			case []interface{}:
				WriteBodyBodyHcl(blockBody, k, vv)
			default:
				blockBody.SetAttributeRaw(k, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", vv))))
			}
		}
	case []interface{}:
		for _, v := range vt {
			WriteAttribute(blockBody, v)
		}
	default:
	}
}
