// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package supportauthzsupportpermit


type SupportauthzSupportPermitPermitActions struct {
	// An explicit list of actions to grant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/supportauthz_support_permit#actions SupportauthzSupportPermit#actions}
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// Grants all actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/supportauthz_support_permit#all_actions SupportauthzSupportPermit#all_actions}
	AllActions *string `field:"optional" json:"allActions" yaml:"allActions"`
}

