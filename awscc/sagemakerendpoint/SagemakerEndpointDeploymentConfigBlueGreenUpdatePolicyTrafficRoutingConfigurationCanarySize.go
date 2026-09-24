// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint


type SagemakerEndpointDeploymentConfigBlueGreenUpdatePolicyTrafficRoutingConfigurationCanarySize struct {
	// Specifies whether the `Value` is an instance count or a capacity unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#type SagemakerEndpoint#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value representing either the number of instances or the number of capacity units.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint#value SagemakerEndpoint#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

