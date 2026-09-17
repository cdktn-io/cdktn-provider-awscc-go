// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaymethod


type ApigatewayMethodMethodResponses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_method#response_models ApigatewayMethod#response_models}.
	ResponseModels *map[string]*string `field:"optional" json:"responseModels" yaml:"responseModels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_method#response_parameters ApigatewayMethod#response_parameters}.
	ResponseParameters *map[string]interface{} `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_method#status_code ApigatewayMethod#status_code}.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

