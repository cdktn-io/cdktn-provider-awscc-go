// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint


type Ec2ClientVpnEndpointAuthenticationOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#type Ec2ClientVpnEndpoint#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#active_directory Ec2ClientVpnEndpoint#active_directory}.
	ActiveDirectory *Ec2ClientVpnEndpointAuthenticationOptionsActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#federated_authentication Ec2ClientVpnEndpoint#federated_authentication}.
	FederatedAuthentication *Ec2ClientVpnEndpointAuthenticationOptionsFederatedAuthentication `field:"optional" json:"federatedAuthentication" yaml:"federatedAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#mutual_authentication Ec2ClientVpnEndpoint#mutual_authentication}.
	MutualAuthentication *Ec2ClientVpnEndpointAuthenticationOptionsMutualAuthentication `field:"optional" json:"mutualAuthentication" yaml:"mutualAuthentication"`
}

