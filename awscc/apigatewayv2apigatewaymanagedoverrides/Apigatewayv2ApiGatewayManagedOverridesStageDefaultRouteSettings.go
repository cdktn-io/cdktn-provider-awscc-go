// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2apigatewaymanagedoverrides


type Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettings struct {
	// Specifies whether detailed metrics are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#detailed_metrics_enabled Apigatewayv2ApiGatewayManagedOverrides#detailed_metrics_enabled}
	DetailedMetricsEnabled interface{} `field:"optional" json:"detailedMetricsEnabled" yaml:"detailedMetricsEnabled"`
	// Specifies the throttling burst limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#throttling_burst_limit Apigatewayv2ApiGatewayManagedOverrides#throttling_burst_limit}
	ThrottlingBurstLimit *float64 `field:"optional" json:"throttlingBurstLimit" yaml:"throttlingBurstLimit"`
	// Specifies the throttling rate limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#throttling_rate_limit Apigatewayv2ApiGatewayManagedOverrides#throttling_rate_limit}
	ThrottlingRateLimit *float64 `field:"optional" json:"throttlingRateLimit" yaml:"throttlingRateLimit"`
}

