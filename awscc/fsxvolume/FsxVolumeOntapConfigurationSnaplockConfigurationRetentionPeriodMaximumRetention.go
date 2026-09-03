// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMaximumRetention struct {
	// Defines the type of time for the retention period of an FSx for ONTAP SnapLock volume.
	//
	// Set it to one of the valid types. If you set it to INFINITE, the files are retained forever. If you set it to UNSPECIFIED, the files are retained until you set an explicit retention period.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#type FsxVolume#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Defines the amount of time for the retention period of an FSx for ONTAP SnapLock volume.
	//
	// You can't set a value for INFINITE or UNSPECIFIED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#value FsxVolume#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

