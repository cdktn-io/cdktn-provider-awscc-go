// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluetable


type GlueTableTableInputStorageDescriptorSkewedInfo struct {
	// A list of names of columns that contain skewed values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#skewed_column_names GlueTable#skewed_column_names}
	SkewedColumnNames *[]*string `field:"optional" json:"skewedColumnNames" yaml:"skewedColumnNames"`
	// A mapping of skewed values to the columns that contain them.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#skewed_column_value_location_maps GlueTable#skewed_column_value_location_maps}
	SkewedColumnValueLocationMaps *string `field:"optional" json:"skewedColumnValueLocationMaps" yaml:"skewedColumnValueLocationMaps"`
	// A list of names of columns that contain skewed values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/glue_table#skewed_column_values GlueTable#skewed_column_values}
	SkewedColumnValues *[]*string `field:"optional" json:"skewedColumnValues" yaml:"skewedColumnValues"`
}

