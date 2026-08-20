// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayapikey


type ApigatewayApiKeyTags struct {
	// A string you can use to assign a value.
	//
	// The combination of tag keys and values can help you organize and categorize your resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_api_key#key ApigatewayApiKey#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the specified tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_api_key#value ApigatewayApiKey#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

