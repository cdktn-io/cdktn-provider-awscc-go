// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationAggregateConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fsx_volume#aggregates FsxVolume#aggregates}.
	Aggregates *[]*string `field:"optional" json:"aggregates" yaml:"aggregates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/fsx_volume#constituents_per_aggregate FsxVolume#constituents_per_aggregate}.
	ConstituentsPerAggregate *float64 `field:"optional" json:"constituentsPerAggregate" yaml:"constituentsPerAggregate"`
}

