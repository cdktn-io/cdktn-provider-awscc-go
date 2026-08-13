// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package nimblestudiolaunchprofile


type NimblestudioLaunchProfileStreamConfigurationSessionStorageRoot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/nimblestudio_launch_profile#linux NimblestudioLaunchProfile#linux}.
	Linux *string `field:"optional" json:"linux" yaml:"linux"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/nimblestudio_launch_profile#windows NimblestudioLaunchProfile#windows}.
	Windows *string `field:"optional" json:"windows" yaml:"windows"`
}

