// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRule struct {
	// Information about the retention period for the snapshot archiving rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#retain_rule DlmLifecyclePolicy#retain_rule}
	RetainRule *DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRuleRetainRule `field:"optional" json:"retainRule" yaml:"retainRule"`
}

