// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsEventSourceParameters struct {
	// The snapshot description that can trigger the policy.
	//
	// The description pattern is specified using a regular expression. The policy runs only if a snapshot with a description that matches the specified pattern is shared with your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#description_regex DlmLifecyclePolicy#description_regex}
	DescriptionRegex *string `field:"optional" json:"descriptionRegex" yaml:"descriptionRegex"`
	// The type of event. Currently, only snapshot sharing events are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#event_type DlmLifecyclePolicy#event_type}
	EventType *string `field:"optional" json:"eventType" yaml:"eventType"`
	// The IDs of the AWS accounts that can trigger policy by sharing snapshots with your account.
	//
	// The policy only runs if one of the specified AWS accounts shares a snapshot with your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dlm_lifecycle_policy#snapshot_owner DlmLifecyclePolicy#snapshot_owner}
	SnapshotOwner *[]*string `field:"optional" json:"snapshotOwner" yaml:"snapshotOwner"`
}

