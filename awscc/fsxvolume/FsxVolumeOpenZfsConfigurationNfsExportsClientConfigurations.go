// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOpenZfsConfigurationNfsExportsClientConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#clients FsxVolume#clients}.
	Clients *string `field:"optional" json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#options FsxVolume#options}.
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
}

