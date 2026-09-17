// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaygatewayresponse

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayGatewayResponseConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_gateway_response#response_type ApigatewayGatewayResponse#response_type}.
	ResponseType *string `field:"required" json:"responseType" yaml:"responseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_gateway_response#rest_api_id ApigatewayGatewayResponse#rest_api_id}.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_gateway_response#response_parameters ApigatewayGatewayResponse#response_parameters}.
	ResponseParameters *map[string]*string `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_gateway_response#response_templates ApigatewayGatewayResponse#response_templates}.
	ResponseTemplates *map[string]*string `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_gateway_response#status_code ApigatewayGatewayResponse#status_code}.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

