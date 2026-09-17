// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRuleRetainRule struct {
	// Information about retention period in the Amazon EBS Snapshots Archive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#retention_archive_tier DlmLifecyclePolicy#retention_archive_tier}
	RetentionArchiveTier *DlmLifecyclePolicyPolicyDetailsSchedulesArchiveRuleRetainRuleRetentionArchiveTier `field:"optional" json:"retentionArchiveTier" yaml:"retentionArchiveTier"`
}

