// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2ipamscope


type Ec2IpamScopeExternalAuthorityConfiguration struct {
	// Resource identifier of the scope in the external service connecting to your AWS IPAM scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_scope#external_resource_identifier Ec2IpamScope#external_resource_identifier}
	ExternalResourceIdentifier *string `field:"optional" json:"externalResourceIdentifier" yaml:"externalResourceIdentifier"`
	// An external service connecting to your AWS IPAM scope.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/ec2_ipam_scope#ipam_scope_external_authority_type Ec2IpamScope#ipam_scope_external_authority_type}
	IpamScopeExternalAuthorityType *string `field:"optional" json:"ipamScopeExternalAuthorityType" yaml:"ipamScopeExternalAuthorityType"`
}

