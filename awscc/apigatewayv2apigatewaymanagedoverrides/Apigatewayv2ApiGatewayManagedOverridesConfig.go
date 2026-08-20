// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2apigatewaymanagedoverrides

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Apigatewayv2ApiGatewayManagedOverridesConfig struct {
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
	// The ID of the API for which to override the configuration of API Gateway-managed resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#api_id Apigatewayv2ApiGatewayManagedOverrides#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// Overrides the integration configuration for an API Gateway-managed integration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#integration Apigatewayv2ApiGatewayManagedOverrides#integration}
	Integration *Apigatewayv2ApiGatewayManagedOverridesIntegration `field:"optional" json:"integration" yaml:"integration"`
	// Overrides the route configuration for an API Gateway-managed route.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#route Apigatewayv2ApiGatewayManagedOverrides#route}
	Route *Apigatewayv2ApiGatewayManagedOverridesRoute `field:"optional" json:"route" yaml:"route"`
	// Overrides the stage configuration for an API Gateway-managed stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigatewayv2_api_gateway_managed_overrides#stage Apigatewayv2ApiGatewayManagedOverrides#stage}
	Stage *Apigatewayv2ApiGatewayManagedOverridesStage `field:"optional" json:"stage" yaml:"stage"`
}

