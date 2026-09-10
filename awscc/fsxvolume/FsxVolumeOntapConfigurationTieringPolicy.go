// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationTieringPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#cooling_period FsxVolume#cooling_period}.
	CoolingPeriod *float64 `field:"optional" json:"coolingPeriod" yaml:"coolingPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_volume#name FsxVolume#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

