// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriod struct {
	// The retention period assigned to a write once, read many (WORM) file by default if an explicit retention period is not set for an FSx for ONTAP SnapLock volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#default_retention FsxVolume#default_retention}
	DefaultRetention *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodDefaultRetention `field:"optional" json:"defaultRetention" yaml:"defaultRetention"`
	// The longest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#maximum_retention FsxVolume#maximum_retention}
	MaximumRetention *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMaximumRetention `field:"optional" json:"maximumRetention" yaml:"maximumRetention"`
	// The shortest retention period that can be assigned to a WORM file on an FSx for ONTAP SnapLock volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#minimum_retention FsxVolume#minimum_retention}
	MinimumRetention *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMinimumRetention `field:"optional" json:"minimumRetention" yaml:"minimumRetention"`
}

