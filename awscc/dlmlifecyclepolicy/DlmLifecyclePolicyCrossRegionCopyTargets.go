// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyCrossRegionCopyTargets struct {
	// The target Region, for example `us-east-1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#target_region DlmLifecyclePolicy#target_region}
	TargetRegion *string `field:"optional" json:"targetRegion" yaml:"targetRegion"`
}

