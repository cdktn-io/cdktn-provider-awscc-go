// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreruntime


type BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpoint struct {
	// Managed VPC resource configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#managed_vpc_resource BedrockagentcoreRuntime#managed_vpc_resource}
	ManagedVpcResource *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointManagedVpcResource `field:"optional" json:"managedVpcResource" yaml:"managedVpcResource"`
	// Self-managed VPC Lattice resource configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrockagentcore_runtime#self_managed_lattice_resource BedrockagentcoreRuntime#self_managed_lattice_resource}
	SelfManagedLatticeResource *BedrockagentcoreRuntimeAuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointOverridesPrivateEndpointSelfManagedLatticeResource `field:"optional" json:"selfManagedLatticeResource" yaml:"selfManagedLatticeResource"`
}

