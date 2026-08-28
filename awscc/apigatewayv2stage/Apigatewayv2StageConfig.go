// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2stage

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type Apigatewayv2StageConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#api_id Apigatewayv2Stage#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The stage name.
	//
	// Stage names can contain only alphanumeric characters, hyphens, and underscores, or be $default. Maximum length is 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#stage_name Apigatewayv2Stage#stage_name}
	StageName *string `field:"required" json:"stageName" yaml:"stageName"`
	// Settings for logging access in this stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#access_log_settings Apigatewayv2Stage#access_log_settings}
	AccessLogSettings *Apigatewayv2StageAccessLogSettings `field:"optional" json:"accessLogSettings" yaml:"accessLogSettings"`
	// Specifies whether updates to an API automatically trigger a new deployment. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#auto_deploy Apigatewayv2Stage#auto_deploy}
	AutoDeploy interface{} `field:"optional" json:"autoDeploy" yaml:"autoDeploy"`
	// The identifier of a client certificate for a Stage. Supported only for WebSocket APIs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#client_certificate_id Apigatewayv2Stage#client_certificate_id}
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// The default route settings for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#default_route_settings Apigatewayv2Stage#default_route_settings}
	DefaultRouteSettings *Apigatewayv2StageDefaultRouteSettings `field:"optional" json:"defaultRouteSettings" yaml:"defaultRouteSettings"`
	// The deployment identifier for the API stage. Can't be updated if autoDeploy is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#deployment_id Apigatewayv2Stage#deployment_id}
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// The description for the API stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#description Apigatewayv2Stage#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Route settings for the stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#route_settings Apigatewayv2Stage#route_settings}
	RouteSettings *string `field:"optional" json:"routeSettings" yaml:"routeSettings"`
	// A map that defines the stage variables for a Stage.
	//
	// Variable names can have alphanumeric and underscore characters, and the values must match [A-Za-z0-9-._~:/?#&=,]+.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#stage_variables Apigatewayv2Stage#stage_variables}
	StageVariables *string `field:"optional" json:"stageVariables" yaml:"stageVariables"`
	// The collection of tags. Each tag element is associated with a given resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_stage#tags Apigatewayv2Stage#tags}
	Tags *string `field:"optional" json:"tags" yaml:"tags"`
}

