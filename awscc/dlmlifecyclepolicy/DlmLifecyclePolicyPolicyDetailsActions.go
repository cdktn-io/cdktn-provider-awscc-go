// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsActions struct {
	// The rule for copying shared snapshots across Regions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#cross_region_copy DlmLifecyclePolicy#cross_region_copy}
	CrossRegionCopy interface{} `field:"optional" json:"crossRegionCopy" yaml:"crossRegionCopy"`
	// A descriptive name for the action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#name DlmLifecyclePolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

