package hcl

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/zclconf/go-cty/cty"
)

func ShowStateFileRaw(state *tfjson.State) {
	file := hclwrite.NewEmptyFile()
	body := file.Body()

	writeBodyHcl(body, state.Values.RootModule.Resources)
	for _, child := range state.Values.RootModule.ChildModules {
		for _, grandChild := range child.ChildModules {
			writeBodyHcl(body, grandChild.Resources)
		}
		writeBodyHcl(body, child.Resources)
	}
	fmt.Println(string(file.Bytes()))
}

func writeBodyHcl(body *hclwrite.Body, resources []*tfjson.StateResource) {
	for _, v := range resources {
		// write comment
		commentToken := hclwrite.Token{
			Type:  hclsyntax.TokenComment,
			Bytes: []byte(fmt.Sprintf("# %v", v.Address)),
		}
		body.AppendUnstructuredTokens(hclwrite.Tokens{&commentToken})
		body.AppendNewline()
		// write resource
		block := body.AppendNewBlock("resource", []string{v.Type, v.Name})
		blockBody := block.Body()
		for k, attr := range v.AttributeValues {
			if attr == nil {
				continue
			}
			switch v := attr.(type) {
			case []interface{}:
				writeBodyBodyHcl(blockBody, k, v)
			case map[string]interface{}:
				if len(v) < 1 {
					// empty objects
					blockBody.SetAttributeRaw(k, hclwrite.Tokens{
						&hclwrite.Token{
							Bytes: []byte(fmt.Sprintf("%v", "{}")),
							Type:  hclsyntax.TokenAnd,
						},
					})
				} else {
					/*
						sample output
						tags_all = {
						    aaa = "bbb"
						}
					*/
					for kk, vv := range v {
						blockBody.SetAttributeValue(k, cty.ObjectVal(map[string]cty.Value{
							kk: cty.StringVal(fmt.Sprintf("%v", vv)),
						}))
					}
				}

			default:
				if v != "" {
					blockBody.SetAttributeRaw(k, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", v))))
				}
			}

		}
		body.AppendNewline()
	}
}

// write block
func writeBodyBodyHcl(body *hclwrite.Body, key string, values []interface{}) {
	if len(values) < 1 {
		return
	}

	block := body.AppendNewBlock(key, []string{})
	blockBody := block.Body()
	for _, value := range values {
		writeAttribute(blockBody, value)
	}

}

// write attribute
func writeAttribute(blockBody *hclwrite.Body, value interface{}) {
	switch vt := value.(type) {
	case map[string]interface{}:
		for k, v := range vt {
			switch vv := v.(type) {
			case []interface{}:
				writeBodyBodyHcl(blockBody, k, vv)
			default:
				blockBody.SetAttributeRaw(k, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", vv))))
			}
		}
	case []interface{}:
		for _, v := range vt {
			writeAttribute(blockBody, v)
		}
	default:
	}

}

func ShowStateFileJson(state *tfjson.State) {
	jsonByte, err := json.Marshal(state)
	if err != nil {
		log.Printf("error convert to json %s", err)

	}
	fmt.Println(string(jsonByte))
}
