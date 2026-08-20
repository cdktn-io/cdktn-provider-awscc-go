// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaydeployment

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ApigatewayDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#rest_api_id ApigatewayDeployment#rest_api_id}.
	RestApiId *string `field:"required" json:"restApiId" yaml:"restApiId"`
	// The ``DeploymentCanarySettings`` property type specifies settings for the canary deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#deployment_canary_settings ApigatewayDeployment#deployment_canary_settings}
	DeploymentCanarySettings *ApigatewayDeploymentDeploymentCanarySettings `field:"optional" json:"deploymentCanarySettings" yaml:"deploymentCanarySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#description ApigatewayDeployment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The description of the Stage resource for the Deployment resource to create.
	//
	// To specify a stage description, you must also provide a stage name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#stage_description ApigatewayDeployment#stage_description}
	StageDescription *ApigatewayDeploymentStageDescription `field:"optional" json:"stageDescription" yaml:"stageDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/apigateway_deployment#stage_name ApigatewayDeployment#stage_name}.
	StageName *string `field:"optional" json:"stageName" yaml:"stageName"`
}

