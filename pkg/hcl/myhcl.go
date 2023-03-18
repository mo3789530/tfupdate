package hcl

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"

	tfjson "github.com/hashicorp/terraform-json"
)

func HclToJson(filepath string) (string, error) {
	jsonObject, err := HclToObject(filepath)
	if err != nil {
		return "", err
	}
	out, err := json.Marshal(jsonObject)
	if err != nil {
		log.Printf("error marshal %s \n", err)
		return "", err
	}

	return string(out), nil
}

func HclToJsonString(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("error read file %s \n", err)
		return "", err
	}
	log.Printf(filepath)
	log.Print(string(data))

	// HCLからJSONに変換
	result, diags := hclsyntax.ParseConfig(data, filepath, hcl.Pos{})
	if diags.HasErrors() {
		panic(diags)
	}
	var jsonObject map[string]interface{}
	// var jsonObject interface{}
	diagnostics := gohcl.DecodeBody(result.Body, nil, &jsonObject)
	if diagnostics.HasErrors() {
		return "", diagnostics.Errs()[0]
	}

	out, err := json.Marshal(jsonObject)
	if err != nil {
		log.Printf("error marshal %s \n", err)
		return "", err
	}

	return string(out), nil

	// 	hclStr := `
	// 	param1 = "value1"
	// 	param2 = 123
	// `

	// 	var config map[string]interface{}
	// 	err := hclsimple.Decode("example.hcl", []byte(hclStr), nil, &config)
	// 	if err != nil {
	// 		fmt.Println("Error decoding HCL:", err)
	// 		return "", nil
	// 	}

	// 	jsonBytes, err := json.Marshal(config)
	// 	if err != nil {
	// 		fmt.Println("Error encoding JSON:", err)
	// 		return "", nil
	// 	}

	// jsonStr := string(jsonBytes)
	// fmt.Println(jsonStr)
	// return "", nil
}

func HclToObject(filepath string) (interface{}, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("error read file %s \n", err)
		return nil, err
	}

	var jsonObject interface{}
	err = hclsimple.Decode("example.", data, nil, &jsonObject)
	if err != nil {
		log.Printf("error decode %s \n", err)
		return nil, nil
	}

	return jsonObject, nil
}

// func JsonToHcl(filepath string) (string, error) {
// 	data, err := os.ReadFile(filepath)
// 	if err != nil {
// 		log.Printf("error read file %s \n", err)
// 		return "", err
// 	}

// 	var obj interface{}
// 	err = json.Unmarshal(data, &obj)
// 	if err != nil {
// 		log.Printf("error unmarshal file %s \n", err)
// 		return "", err
// 	}

// 	// convert to hcl
// 	f := hclwrite.NewEmptyFile()
// 	state, err := objToTfState(string(data))
// 	if err != nil {
// 		log.Printf(err.Error())
// 	}
// 	state.Values.RootModule.Resources

//		gohcl.EncodeIntoBody(obj, f.Body())
//		return string(f.Bytes()), nil
//	}

func JsonToHcl(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data map[string]interface{}
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		panic(err)
	}

	f := hclwrite.NewEmptyFile()
	body := f.Body()
	for k, v := range data {
		block := body.AppendNewBlock(k, []string{})
		encodeBlock(block.Body(), v)
	}

	fmt.Println(string(f.Bytes()))
	return "", nil
}

func encodeBlock(body *hclwrite.Body, data interface{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		for k, vv := range v {
			block := body.AppendNewBlock(k, []string{})
			encodeBlock(block.Body(), vv)
		}
	case []interface{}:
		for _, vv := range v {
			encodeBlock(body, vv)
		}
	default:
		body.SetAttributeValue("value", cty.StringVal(fmt.Sprintf("%v", v)))
	}
}
func objToTfState(obj string) (*tfjson.State, error) {
	fmt.Print(obj)
	var stateReader io.Reader = strings.NewReader(obj)
	var state *tfjson.State

	if err := json.NewDecoder(stateReader).Decode(&state); err != nil {
		log.Printf("error %s", err)
		return nil, err
	}

	return state, nil
}
