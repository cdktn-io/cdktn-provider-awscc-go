// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sestenant


type SesTenantSuppressionAttributes struct {
	// The reasons that cause messages to be added to the suppression list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/ses_tenant#suppressed_reasons SesTenant#suppressed_reasons}
	SuppressedReasons *[]*string `field:"optional" json:"suppressedReasons" yaml:"suppressedReasons"`
	// The scope of the suppression list for the tenant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/ses_tenant#suppression_scope SesTenant#suppression_scope}
	SuppressionScope *string `field:"optional" json:"suppressionScope" yaml:"suppressionScope"`
}

