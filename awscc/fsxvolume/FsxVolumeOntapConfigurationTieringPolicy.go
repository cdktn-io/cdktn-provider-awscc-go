// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationTieringPolicy struct {
	// Specifies the number of days that user data in a volume must remain inactive before it is considered "cold" and moved to the capacity pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#cooling_period FsxVolume#cooling_period}
	CoolingPeriod *float64 `field:"optional" json:"coolingPeriod" yaml:"coolingPeriod"`
	// Specifies the tiering policy used to transition data. Default value is SNAPSHOT_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#name FsxVolume#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
}

