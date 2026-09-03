// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supportauthzsupportpermit


type SupportauthzSupportPermitPermitConditions struct {
	// The permit is active only after this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/supportauthz_support_permit#allow_after SupportauthzSupportPermit#allow_after}
	AllowAfter *string `field:"optional" json:"allowAfter" yaml:"allowAfter"`
	// The permit is active only before this time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/supportauthz_support_permit#allow_before SupportauthzSupportPermit#allow_before}
	AllowBefore *string `field:"optional" json:"allowBefore" yaml:"allowBefore"`
}

