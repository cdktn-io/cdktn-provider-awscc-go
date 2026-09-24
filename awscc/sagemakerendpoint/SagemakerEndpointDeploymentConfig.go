// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint


type SagemakerEndpointDeploymentConfig struct {
	// Configuration for automatic rollback if an error occurs during deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#auto_rollback_configuration SagemakerEndpoint#auto_rollback_configuration}
	AutoRollbackConfiguration *SagemakerEndpointDeploymentConfigAutoRollbackConfiguration `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// Configuration for blue-green update deployment policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#blue_green_update_policy SagemakerEndpoint#blue_green_update_policy}
	BlueGreenUpdatePolicy *SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicy `field:"optional" json:"blueGreenUpdatePolicy" yaml:"blueGreenUpdatePolicy"`
	// Configuration for rolling update deployment policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#rolling_update_policy SagemakerEndpoint#rolling_update_policy}
	RollingUpdatePolicy *SagemakerEndpointDeploymentConfigRollingUpdatePolicy `field:"optional" json:"rollingUpdatePolicy" yaml:"rollingUpdatePolicy"`
}

