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

	for _, child := range state.Values.RootModule.ChildModules {
		for _, v := range child.ChildModules[0].Resources {
			// write comment
			commentToken := hclwrite.Token{
				Type:  hclsyntax.TokenComment,
				Bytes: []byte(fmt.Sprintf("# %v.%v", v.Type, v.Name)),
			}
			body.AppendUnstructuredTokens(hclwrite.Tokens{&commentToken})
			body.AppendNewline()
			block := body.AppendNewBlock("resources", []string{v.Type, v.Name})
			blockBody := block.Body()
			for k, atter := range v.AttributeValues {
				if atter == nil {
					continue
				}
				blockBody.SetAttributeRaw(k, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", atter))))

			}
			body.AppendNewline()
		}
	}

	for _, v := range state.Values.RootModule.Resources {
		log.Panicln(v.Type)
		// write comment
		commentToken := hclwrite.Token{
			Type:  hclsyntax.TokenComment,
			Bytes: []byte(fmt.Sprintf("# %v.%v", v.Type, v.Name)),
		}
		body.AppendUnstructuredTokens(hclwrite.Tokens{&commentToken})
		body.AppendNewline()
		// block := body.AppendNewBlock("resources", []string{v.Type, v.Name})
		// blockBody := block.Body()
		// for k, atter := range v.AttributeValues {
		// 	if atter == nil {
		// 		continue
		// 	}
		// 	blockBody.SetAttributeRaw(k, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", atter))))

		// }
		body.AppendNewline()
	}
	fmt.Println(string(file.Bytes()))
}

func writeHcl(body *hclwrite.Body, resources []*tfjson.StateResource) error {
	for _, v := range resources {
		log.Panicln(v.Type)
		// write comment
		commentToken := hclwrite.Token{
			Type:  hclsyntax.TokenComment,
			Bytes: []byte(fmt.Sprintf("# %v.%v", v.Type, v.Name)),
		}
		body.AppendUnstructuredTokens(hclwrite.Tokens{&commentToken})
		body.AppendNewline()
		block := body.AppendNewBlock("resources", []string{v.Type, v.Name})
		blockBody := block.Body()
		for k, atter := range v.AttributeValues {
			if atter == nil {
				continue
			}
			blockBody.SetAttributeRaw(k, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", atter))))

		}
		body.AppendNewline()
	}

}

func ShowStateFileJson(state *tfjson.State) {
	jsonByte, err := json.Marshal(state)
	if err != nil {
		log.Printf("error convert to json %s", err)

	}
	fmt.Println(string(jsonByte))
}
