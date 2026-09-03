// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaymethod


type ApigatewayMethodIntegrationIntegrationResponses struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_method#content_handling ApigatewayMethod#content_handling}.
	ContentHandling *string `field:"optional" json:"contentHandling" yaml:"contentHandling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_method#response_parameters ApigatewayMethod#response_parameters}.
	ResponseParameters *map[string]*string `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_method#response_templates ApigatewayMethod#response_templates}.
	ResponseTemplates *map[string]*string `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_method#selection_pattern ApigatewayMethod#selection_pattern}.
	SelectionPattern *string `field:"optional" json:"selectionPattern" yaml:"selectionPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_method#status_code ApigatewayMethod#status_code}.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
}

