// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsProjectOperation struct {
	// <p>Projected columns.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_set#projected_columns QuicksightDataSet#projected_columns}
	ProjectedColumns *[]*string `field:"optional" json:"projectedColumns" yaml:"projectedColumns"`
}

