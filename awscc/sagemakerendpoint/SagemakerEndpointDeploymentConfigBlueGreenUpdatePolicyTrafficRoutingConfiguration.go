// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint


type SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfiguration struct {
	// Specifies the size of the canary traffic in a canary deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#canary_size SagemakerEndpoint#canary_size}
	CanarySize *SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationCanarySize `field:"optional" json:"canarySize" yaml:"canarySize"`
	// Specifies the step size for linear traffic routing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#linear_step_size SagemakerEndpoint#linear_step_size}
	LinearStepSize *SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationLinearStepSize `field:"optional" json:"linearStepSize" yaml:"linearStepSize"`
	// Specifies the type of traffic routing (e.g., 'AllAtOnce', 'Canary', 'Linear').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#type SagemakerEndpoint#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the wait interval between traffic shifts, in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#wait_interval_in_seconds SagemakerEndpoint#wait_interval_in_seconds}
	WaitIntervalInSeconds *float64 `field:"optional" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
}

