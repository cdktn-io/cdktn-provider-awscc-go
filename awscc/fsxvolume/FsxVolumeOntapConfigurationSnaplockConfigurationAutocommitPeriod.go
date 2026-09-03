// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationSnaplockConfigurationAutocommitPeriod struct {
	// Defines the type of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.
	//
	// Setting this value to NONE disables autocommit. The default value is NONE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#type FsxVolume#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Defines the amount of time for the autocommit period of a file in an FSx for ONTAP SnapLock volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#value FsxVolume#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

