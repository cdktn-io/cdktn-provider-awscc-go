// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreoauth2credentialprovider


type BedrockagentcoreOAuth2CredentialProviderOauth2ProviderConfigInputCustomOauth2ProviderConfigPrivateEndpointManagedVpcResource struct {
	// The IP address type for the resource configuration endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#endpoint_ip_address_type BedrockagentcoreOAuth2CredentialProvider#endpoint_ip_address_type}
	EndpointIpAddressType *string `field:"optional" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// An intermediate publicly resolvable domain used as the VPC Lattice resource configuration endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#routing_domain BedrockagentcoreOAuth2CredentialProvider#routing_domain}
	RoutingDomain *string `field:"optional" json:"routingDomain" yaml:"routingDomain"`
	// The security group IDs to associate with the VPC Lattice resource gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#security_group_ids BedrockagentcoreOAuth2CredentialProvider#security_group_ids}
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// The subnet IDs within the VPC where the VPC Lattice resource gateway is placed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#subnet_ids BedrockagentcoreOAuth2CredentialProvider#subnet_ids}
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Tags to apply to the managed VPC Lattice resource gateway.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#tags BedrockagentcoreOAuth2CredentialProvider#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the VPC that contains your private resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/bedrockagentcore_o_auth_2_credential_provider#vpc_identifier BedrockagentcoreOAuth2CredentialProvider#vpc_identifier}
	VpcIdentifier *string `field:"optional" json:"vpcIdentifier" yaml:"vpcIdentifier"`
}

