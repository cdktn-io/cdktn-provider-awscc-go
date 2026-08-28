// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeCapacityProviderConfiguration struct {
	// ARN of the capacity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_runtime#capacity_provider_arn BedrockagentcoreRuntime#capacity_provider_arn}
	CapacityProviderArn *string `field:"optional" json:"capacityProviderArn" yaml:"capacityProviderArn"`
}

