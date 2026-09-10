// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#default_retention FsxVolume#default_retention}.
	DefaultRetention *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodDefaultRetention `field:"optional" json:"defaultRetention" yaml:"defaultRetention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#maximum_retention FsxVolume#maximum_retention}.
	MaximumRetention *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMaximumRetention `field:"optional" json:"maximumRetention" yaml:"maximumRetention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#minimum_retention FsxVolume#minimum_retention}.
	MinimumRetention *FsxVolumeOntapConfigurationSnaplockConfigurationRetentionPeriodMinimumRetention `field:"optional" json:"minimumRetention" yaml:"minimumRetention"`
}

