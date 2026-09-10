// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlylaunch


type EvidentlyLaunchScheduledSplitsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evidently_launch#group_weights EvidentlyLaunch#group_weights}.
	GroupWeights interface{} `field:"required" json:"groupWeights" yaml:"groupWeights"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evidently_launch#start_time EvidentlyLaunch#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/evidently_launch#segment_overrides EvidentlyLaunch#segment_overrides}.
	SegmentOverrides interface{} `field:"optional" json:"segmentOverrides" yaml:"segmentOverrides"`
}

