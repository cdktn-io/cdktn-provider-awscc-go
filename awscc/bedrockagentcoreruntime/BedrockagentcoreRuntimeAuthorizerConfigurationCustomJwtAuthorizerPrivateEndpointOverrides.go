// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverrides struct {
	// The domain to override.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_runtime#domain BedrockagentcoreRuntime#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Private endpoint configuration. Exactly one of SelfManagedLatticeResource or ManagedVpcResource must be specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_runtime#private_endpoint BedrockagentcoreRuntime#private_endpoint}
	PrivateEndpoint *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpoint `field:"optional" json:"privateEndpoint" yaml:"privateEndpoint"`
}

