// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOntapConfigurationAggregateConfiguration struct {
	// The list of aggregates that this volume resides on.
	//
	// Aggregates are storage pools which make up your primary storage tier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#aggregates FsxVolume#aggregates}
	Aggregates *[]*string `field:"optional" json:"aggregates" yaml:"aggregates"`
	// Used to explicitly set the number of constituents within the FlexGroup per storage aggregate.
	//
	// This field is optional when creating a FlexGroup volume. If unspecified, the default value will be 8. This field cannot be provided when creating a FlexVol volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#constituents_per_aggregate FsxVolume#constituents_per_aggregate}
	ConstituentsPerAggregate *float64 `field:"optional" json:"constituentsPerAggregate" yaml:"constituentsPerAggregate"`
}

