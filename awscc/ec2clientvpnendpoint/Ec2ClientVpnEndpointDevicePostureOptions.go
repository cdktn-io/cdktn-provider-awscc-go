// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint


type Ec2ClientVpnEndpointDevicePostureOptions struct {
	// Whether device posture is enabled for the Client VPN endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#enabled Ec2ClientVpnEndpoint#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The device-posture trust providers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#trust_providers Ec2ClientVpnEndpoint#trust_providers}
	TrustProviders interface{} `field:"optional" json:"trustProviders" yaml:"trustProviders"`
}

