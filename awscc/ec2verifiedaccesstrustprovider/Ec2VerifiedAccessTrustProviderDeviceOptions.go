// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccesstrustprovider


type Ec2VerifiedAccessTrustProviderDeviceOptions struct {
	// URL Verified Access will use to verify authenticity of the device tokens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_verified_access_trust_provider#public_signing_key_url Ec2VerifiedAccessTrustProvider#public_signing_key_url}
	PublicSigningKeyUrl *string `field:"optional" json:"publicSigningKeyUrl" yaml:"publicSigningKeyUrl"`
	// The ID of the tenant application with the device-identity provider.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ec2_verified_access_trust_provider#tenant_id Ec2VerifiedAccessTrustProvider#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

