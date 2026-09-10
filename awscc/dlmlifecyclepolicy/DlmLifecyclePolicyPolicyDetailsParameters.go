// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsParameters struct {
	// **[Custom snapshot policies that target instances only]** Indicates whether to exclude the root volume from multi-volume snapshot sets.
	//
	// The default is `false`. If you specify `true`, then the root volumes attached to targeted instances will be excluded from the multi-volume snapshot sets created by the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#exclude_boot_volume DlmLifecyclePolicy#exclude_boot_volume}
	ExcludeBootVolume interface{} `field:"optional" json:"excludeBootVolume" yaml:"excludeBootVolume"`
	// **[Custom snapshot policies that target instances only]** The tags used to identify data (non-root) volumes to exclude from multi-volume snapshot sets.
	//
	// If you create a snapshot lifecycle policy that targets instances and you specify tags for this parameter, then data volumes with the specified tags that are attached to targeted instances will be excluded from the multi-volume snapshot sets created by the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#exclude_data_volume_tags DlmLifecyclePolicy#exclude_data_volume_tags}
	ExcludeDataVolumeTags interface{} `field:"optional" json:"excludeDataVolumeTags" yaml:"excludeDataVolumeTags"`
	// **[Custom AMI policies only]** Indicates whether targeted instances are rebooted when the lifecycle policy runs.
	//
	// `true` indicates that targeted instances are not rebooted when the policy runs. `false` indicates that target instances are rebooted when the policy runs.
	//
	// The default is `true` (instances are not rebooted).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dlm_lifecycle_policy#no_reboot DlmLifecyclePolicy#no_reboot}
	NoReboot interface{} `field:"optional" json:"noReboot" yaml:"noReboot"`
}

