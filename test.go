package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	// JSONを読み込む
	jsonStr := `
	{
		"format_version": "1.0",
		"terraform_version": "1.3.9",
		"values": {
			"root_module": {
				"resources": [
					{
						"address": "aws_s3_bucket.my-bucket",
						"mode": "managed",
						"type": "aws_s3_bucket",
						"name": "my-bucket",
						"provider_name": "registry.terraform.io/hashicorp/aws",
						"schema_version": 0,
						"values": {
							"acceleration_status": "",
							"acl": null,
							"arn": "arn:aws:s3:::my-bucket",
							"bucket": "my-bucket",
							"bucket_domain_name": "my-bucket.s3.amazonaws.com",
							"bucket_prefix": null,
							"bucket_regional_domain_name": "my-bucket.s3.ap-northeast-1.amazonaws.com",
							"cors_rule": [],
							"force_destroy": false,
							"grant": [
								{
									"id": "75aa57f09aa0c8caeab4f8c24e99d10f8e7faeebf76c078efc7c6caea54ba06a",
									"permissions": [
										"FULL_CONTROL"
									],
									"type": "CanonicalUser",
									"uri": ""
								}
							],
							"hosted_zone_id": "Z2M4EHUR26P7ZW",
							"id": "my-bucket",
							"lifecycle_rule": [],
							"logging": [],
							"object_lock_configuration": [],
							"object_lock_enabled": false,
							"policy": "",
							"region": "ap-northeast-1",
							"replication_configuration": [],
							"request_payer": "BucketOwner",
							"server_side_encryption_configuration": [],
							"tags": null,
							"tags_all": {},
							"timeouts": null,
							"versioning": [
								{
									"enabled": false,
									"mfa_delete": false
								}
							],
							"website": [],
							"website_domain": null,
							"website_endpoint": null
						},
						"sensitive_values": {
							"cors_rule": [],
							"grant": [
								{
									"permissions": [
										false
									]
								}
							],
							"lifecycle_rule": [],
							"logging": [],
							"object_lock_configuration": [],
							"replication_configuration": [],
							"server_side_encryption_configuration": [],
							"tags_all": {},
							"versioning": [
								{}
							],
							"website": []
						}
					}
				]
			}
		}
	}
`
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("--------------")
	fmt.Println(data["values"].(map[string]interface{})["root_module"])
	// Terraformのコードを生成する
	f := hclwrite.NewFile()
	rootBody := f.Body()
	for key, value := range data["values"].(map[string]interface{})["root_module"].(map[string]interface{}) {
		fmt.Println("--------------")
		fmt.Println(value)
		resourceBody := rootBody.AppendNewBlock("resource", []string{key})
		resourceBodyBody := resourceBody.Body()
		for resKey, resValue := range value.(map[string]interface{}) {
			fmt.Println(resKey, resValue)
			fmt.Println(resourceBodyBody)
			blockBody := resourceBodyBody.AppendNewBlock(resKey, nil)
			fmt.Println(blockBody)
			fmt.Println("--------")
			for bKey, bValue := range resValue.(map[string]interface{}) {
				fmt.Println(bKey)
				fmt.Println(bValue)
				for cKey, cVabValue := range bValue.(map[string]interface{}) {
					fmt.Println(cKey, cVabValue)
					blockBody.Body().SetAttributeRaw(cKey, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", cVabValue))))
				}
				// 	// blockBody.Body().
				// blockBody.Body().SetAttributeRaw(bKey, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", bValue))))
				// blockBody.SetAttributeRaw(bKey, hclwrite.TokensForValue(cty.StringVal(fmt.Sprintf("%v", bValue))))
			}
		}
	}

	// Terraformのコードを出力する
	fmt.Println(string(f.Bytes()))
}

// func main() {
// 	f := hclwrite.NewEmptyFile()
// 	body := f.Body()
// 	header := hclwrite.NewEmptyFile().Body().AppendNewBlock("resource", []string{"aws_s3_bucket", "my_bucket"})

// 	body.AppendNewline()

// 	// body2 := hclwrite.NewEmptyFile().Body().SetAttributeRaw("key", hclwrite.Tokens{})
// 	header.Body().SetAttributeRaw("key", hclwrite.TokensForValue(cty.StringVal("aaa")))
// 	body.AppendBlock(header)

// 	// body.SetAttributeValue("key", hclwrite.StringVal("my_key"))
// 	// body.SetAttributeValue("value", hclwrite.StringVal("my_value"))

// 	fmt.Println(string(f.Bytes()))
// }
