// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentconfig


type CodedeployDeploymentConfigTrafficRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_config#time_based_canary CodedeployDeploymentConfig#time_based_canary}.
	TimeBasedCanary *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedCanary `field:"optional" json:"timeBasedCanary" yaml:"timeBasedCanary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_config#time_based_linear CodedeployDeploymentConfig#time_based_linear}.
	TimeBasedLinear *CodedeployDeploymentConfigTrafficRoutingConfigTimeBasedLinear `field:"optional" json:"timeBasedLinear" yaml:"timeBasedLinear"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/codedeploy_deployment_config#type CodedeployDeploymentConfig#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

