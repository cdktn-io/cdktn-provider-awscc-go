// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supportauthzsupportpermit


type SupportauthzSupportPermitSigningKeyInfo struct {
	// The ARN of the KMS key used to sign permit grants.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/supportauthz_support_permit#kms_key SupportauthzSupportPermit#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

