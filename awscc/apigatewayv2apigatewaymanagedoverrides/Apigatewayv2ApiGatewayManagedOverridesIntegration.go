// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2apigatewaymanagedoverrides


type Apigatewayv2ApiGatewayManagedOverridesIntegration struct {
	// The description of the integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#description Apigatewayv2ApiGatewayManagedOverrides#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Specifies the integration's HTTP method type.
	//
	// For WebSocket APIs, if you use a Lambda integration, you must set the integration method to POST.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#integration_method Apigatewayv2ApiGatewayManagedOverrides#integration_method}
	IntegrationMethod *string `field:"optional" json:"integrationMethod" yaml:"integrationMethod"`
	// Specifies the format of the payload sent to an integration.
	//
	// Required for HTTP APIs. For HTTP APIs, supported values for Lambda proxy integrations are 1.0 and 2.0. For all other integrations, 1.0 is the only supported value
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#payload_format_version Apigatewayv2ApiGatewayManagedOverrides#payload_format_version}
	PayloadFormatVersion *string `field:"optional" json:"payloadFormatVersion" yaml:"payloadFormatVersion"`
	// Custom timeout between 50 and 29,000 milliseconds for WebSocket APIs and between 50 and 30,000 milliseconds for HTTP APIs.
	//
	// The default timeout is 29 seconds for WebSocket APIs and 30 seconds for HTTP APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#timeout_in_millis Apigatewayv2ApiGatewayManagedOverrides#timeout_in_millis}
	TimeoutInMillis *float64 `field:"optional" json:"timeoutInMillis" yaml:"timeoutInMillis"`
}

