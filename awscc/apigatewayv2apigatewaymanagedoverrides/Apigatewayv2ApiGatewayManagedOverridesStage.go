// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2apigatewaymanagedoverrides


type Apigatewayv2ApiGatewayManagedOverridesStage struct {
	// Settings for logging access in a stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#access_log_settings Apigatewayv2ApiGatewayManagedOverrides#access_log_settings}
	AccessLogSettings *Apigatewayv2ApiGatewayManagedOverridesStageAccessLogSettings `field:"optional" json:"accessLogSettings" yaml:"accessLogSettings"`
	// Specifies whether updates to an API automatically trigger a new deployment. The default value is true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#auto_deploy Apigatewayv2ApiGatewayManagedOverrides#auto_deploy}
	AutoDeploy interface{} `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The default route settings for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#default_route_settings Apigatewayv2ApiGatewayManagedOverrides#default_route_settings}
	DefaultRouteSettings *Apigatewayv2ApiGatewayManagedOverridesStageDefaultRouteSettings `field:"optional" json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// The description for the API stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#description Apigatewayv2ApiGatewayManagedOverrides#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Route settings for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#route_settings Apigatewayv2ApiGatewayManagedOverrides#route_settings}
	RouteSettings interface{} `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// A map that defines the stage variables for a Stage.
	//
	// Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#stage_variables Apigatewayv2ApiGatewayManagedOverrides#stage_variables}
	StageVariables *map[string]*string `field:"optional" json:"stageVariables" yaml:"stageVariables"`
}

