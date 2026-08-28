// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2integrationresponse

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Apigatewayv2IntegrationResponseConfig struct {
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
	// The API identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_integration_response#api_id Apigatewayv2IntegrationResponse#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The integration ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_integration_response#integration_id Apigatewayv2IntegrationResponse#integration_id}
	IntegrationId *string `field:"required" json:"integrationId" yaml:"integrationId"`
	// The integration response key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_integration_response#integration_response_key Apigatewayv2IntegrationResponse#integration_response_key}
	IntegrationResponseKey *string `field:"required" json:"integrationResponseKey" yaml:"integrationResponseKey"`
	// Supported only for WebSocket APIs.
	//
	// Specifies how to handle response payload content type conversions. Supported values are ``CONVERT_TO_BINARY`` and ``CONVERT_TO_TEXT``, with the following behaviors:
	//  ``CONVERT_TO_BINARY``: Converts a response payload from a Base64-encoded string to the corresponding binary blob.
	//  ``CONVERT_TO_TEXT``: Converts a response payload from a binary blob to a Base64-encoded string.
	//  If this property is not defined, the response payload will be passed through from the integration response to the route response or method response without modification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_integration_response#content_handling_strategy Apigatewayv2IntegrationResponse#content_handling_strategy}
	ContentHandlingStrategy *string `field:"optional" json:"contentHandlingStrategy" yaml:"contentHandlingStrategy"`
	// A key-value map specifying response parameters that are passed to the method response from the backend.
	//
	// The key is a method response header parameter name and the mapped value is an integration response header value, a static value enclosed within a pair of single quotes, or a JSON expression from the integration response body. The mapping key must match the pattern of ``method.response.header.{name}``, where name is a valid and unique header name. The mapped non-static value must match the pattern of ``integration.response.header.{name}`` or ``integration.response.body.{JSON-expression}``, where ``{name}`` is a valid and unique response header name and ``{JSON-expression}`` is a valid JSON expression without the ``$`` prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_integration_response#response_parameters Apigatewayv2IntegrationResponse#response_parameters}
	ResponseParameters *string `field:"optional" json:"responseParameters" yaml:"responseParameters"`
	// The collection of response templates for the integration response as a string-to-string map of key-value pairs.
	//
	// Response templates are represented as a key/value map, with a content-type as the key and a template as the value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_integration_response#response_templates Apigatewayv2IntegrationResponse#response_templates}
	ResponseTemplates *string `field:"optional" json:"responseTemplates" yaml:"responseTemplates"`
	// The template selection expression for the integration response. Supported only for WebSocket APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_integration_response#template_selection_expression Apigatewayv2IntegrationResponse#template_selection_expression}
	TemplateSelectionExpression *string `field:"optional" json:"templateSelectionExpression" yaml:"templateSelectionExpression"`
}

