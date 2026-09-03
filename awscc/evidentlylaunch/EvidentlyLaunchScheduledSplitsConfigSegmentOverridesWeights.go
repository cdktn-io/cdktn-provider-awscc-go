// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlylaunch


type EvidentlyLaunchScheduledSplitsConfigSegmentOverridesWeights struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evidently_launch#group_name EvidentlyLaunch#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/evidently_launch#split_weight EvidentlyLaunch#split_weight}.
	SplitWeight *float64 `field:"optional" json:"splitWeight" yaml:"splitWeight"`
}

