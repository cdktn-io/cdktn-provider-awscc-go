// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2clientvpnendpoint


type Ec2ClientVpnEndpointDevicePostureOptionsTrustProviders struct {
	// The URL of the public signing key.
	//
	// Conditionally required by the provider type and enforced server-side (crowdstrike/jumpcloud require it, jamf rejects it).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#public_signing_key_url Ec2ClientVpnEndpoint#public_signing_key_url}
	PublicSigningKeyUrl *string `field:"optional" json:"publicSigningKeyUrl" yaml:"publicSigningKeyUrl"`
	// The tenant identifier registered with the trust provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#tenant_id Ec2ClientVpnEndpoint#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// The device-posture trust provider type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/ec2_client_vpn_endpoint#trust_provider_type Ec2ClientVpnEndpoint#trust_provider_type}
	TrustProviderType *string `field:"optional" json:"trustProviderType" yaml:"trustProviderType"`
}

