// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesemailidentity


type SesEmailIdentityDkimSigningAttributes struct {
	// [Bring Your Own DKIM] A private key that's used to generate a DKIM signature.
	//
	// The private key must use 1024 or 2048-bit RSA encryption, and must be encoded using base64 encoding.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_email_identity#domain_signing_private_key SesEmailIdentity#domain_signing_private_key}
	DomainSigningPrivateKey *string `field:"optional" json:"domainSigningPrivateKey" yaml:"domainSigningPrivateKey"`
	// [Bring Your Own DKIM] A string that's used to identify a public key in the DNS configuration for a domain.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_email_identity#domain_signing_selector SesEmailIdentity#domain_signing_selector}
	DomainSigningSelector *string `field:"optional" json:"domainSigningSelector" yaml:"domainSigningSelector"`
	// [Easy DKIM] The key length of the future DKIM key pair to be generated.
	//
	// This can be changed at most once per day.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_email_identity#next_signing_key_length SesEmailIdentity#next_signing_key_length}
	NextSigningKeyLength *string `field:"optional" json:"nextSigningKeyLength" yaml:"nextSigningKeyLength"`
}

