// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint


type Ec2ClientVpnEndpointAuthorizationPolicy struct {
	// A brief description of the authorization policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#description Ec2ClientVpnEndpoint#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Cedar policy document.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#policy_document Ec2ClientVpnEndpoint#policy_document}
	PolicyDocument *string `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// Whether the authorization policy runs in shadow (observe-only) mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#shadow_mode Ec2ClientVpnEndpoint#shadow_mode}
	ShadowMode *string `field:"optional" json:"shadowMode" yaml:"shadowMode"`
}

