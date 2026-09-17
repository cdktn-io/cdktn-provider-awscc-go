// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperation struct {
	// <p>The name of the column to be renamed.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// <p>The new name for the column.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_set#new_column_name QuicksightDataSet#new_column_name}
	NewColumnName *string `field:"optional" json:"newColumnName" yaml:"newColumnName"`
}

