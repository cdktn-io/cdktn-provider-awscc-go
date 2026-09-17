// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperation struct {
	// <p>Column name.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// <p>When casting a column from string to datetime type, you can supply a string in a             format supported by Amazon QuickSight to denote the source data format.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_set#format QuicksightDataSet#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/quicksight_data_set#new_column_type QuicksightDataSet#new_column_type}.
	NewColumnType *string `field:"optional" json:"newColumnType" yaml:"newColumnType"`
}

