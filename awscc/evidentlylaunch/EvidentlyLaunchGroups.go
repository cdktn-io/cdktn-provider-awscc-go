// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package evidentlylaunch


type EvidentlyLaunchGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evidently_launch#feature EvidentlyLaunch#feature}.
	Feature *string `field:"required" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evidently_launch#group_name EvidentlyLaunch#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evidently_launch#variation EvidentlyLaunch#variation}.
	Variation *string `field:"required" json:"variation" yaml:"variation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/evidently_launch#description EvidentlyLaunch#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

