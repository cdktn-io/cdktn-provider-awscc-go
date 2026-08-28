// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateEndpointOverridesPrivateEndpoint struct {
	// Configuration for a managed VPC Lattice resource.
	//
	// AgentCore creates and manages the VPC Lattice resource gateway and resource configuration on your behalf.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#managed_vpc_resource BedrockagentcoreOAuth2CredentialProvider#managed_vpc_resource}
	ManagedVpcResource *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateEndpointOverridesPrivateEndpointManagedVpcResource `field:"optional" json:"managedVpcResource" yaml:"managedVpcResource"`
	// Configuration for a self-managed VPC Lattice resource.
	//
	// You create and manage the VPC Lattice resource gateway and resource configuration, then provide the resource configuration identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#self_managed_lattice_resource BedrockagentcoreOAuth2CredentialProvider#self_managed_lattice_resource}
	SelfManagedLatticeResource *BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateEndpointOverridesPrivateEndpointSelfManagedLatticeResource `field:"optional" json:"selfManagedLatticeResource" yaml:"selfManagedLatticeResource"`
}

