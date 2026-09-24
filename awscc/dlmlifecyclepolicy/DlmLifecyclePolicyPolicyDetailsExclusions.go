// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dlmlifecyclepolicy


type DlmLifecyclePolicyPolicyDetailsExclusions struct {
	// **[Default policies for EBS snapshots only]** Indicates whether to exclude volumes that are attached to instances as the boot volume.
	//
	// If you exclude boot volumes, only volumes attached as data (non-boot) volumes will be backed up by the policy. To exclude boot volumes, specify `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#exclude_boot_volumes DlmLifecyclePolicy#exclude_boot_volumes}
	ExcludeBootVolumes interface{} `field:"optional" json:"excludeBootVolumes" yaml:"excludeBootVolumes"`
	// **[Default policies for EBS-backed AMIs only]** Specifies whether to exclude volumes that have specific tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#exclude_tags DlmLifecyclePolicy#exclude_tags}
	ExcludeTags interface{} `field:"optional" json:"excludeTags" yaml:"excludeTags"`
	// **[Default policies for EBS snapshots only]** Specifies the volume types to exclude.
	//
	// Volumes of the specified types will not be targeted by the policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dlm_lifecycle_policy#exclude_volume_types DlmLifecyclePolicy#exclude_volume_types}
	ExcludeVolumeTypes *[]*string `field:"optional" json:"excludeVolumeTypes" yaml:"excludeVolumeTypes"`
}

