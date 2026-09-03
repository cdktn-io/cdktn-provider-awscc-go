// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supportauthzsupportpermit


type SupportauthzSupportPermitPermit struct {
	// The set of actions a support permit grants. Exactly one of AllActions or Actions must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/supportauthz_support_permit#actions SupportauthzSupportPermit#actions}
	Actions *SupportauthzSupportPermitPermitActions `field:"required" json:"actions" yaml:"actions"`
	// The set of resources a support permit applies to. Exactly one of AllResourcesInRegion or Resources must be provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/supportauthz_support_permit#resources SupportauthzSupportPermit#resources}
	Resources *SupportauthzSupportPermitPermitResources `field:"required" json:"resources" yaml:"resources"`
	// Optional time-bound conditions (at most two).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/supportauthz_support_permit#conditions SupportauthzSupportPermit#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
}

