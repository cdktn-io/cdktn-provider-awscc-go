// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransforms struct {
	// <p>A transform operation that casts a column to a different type.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#cast_column_type_operation QuicksightDataSet#cast_column_type_operation}
	CastColumnTypeOperation *QuicksightDataSetLogicalTableMapDataTransformsCastColumnTypeOperation `field:"optional" json:"castColumnTypeOperation" yaml:"castColumnTypeOperation"`
	// <p>A transform operation that creates calculated columns.
	//
	// Columns created in one such
	//             operation form a lexical closure.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#create_columns_operation QuicksightDataSet#create_columns_operation}
	CreateColumnsOperation *QuicksightDataSetLogicalTableMapDataTransformsCreateColumnsOperation `field:"optional" json:"createColumnsOperation" yaml:"createColumnsOperation"`
	// <p>A transform operation that filters rows based on a condition.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#filter_operation QuicksightDataSet#filter_operation}
	FilterOperation *QuicksightDataSetLogicalTableMapDataTransformsFilterOperation `field:"optional" json:"filterOperation" yaml:"filterOperation"`
	// <p>A transform operation that projects columns.
	//
	// Operations that come after a projection
	//             can only refer to projected columns.</p>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#project_operation QuicksightDataSet#project_operation}
	ProjectOperation *QuicksightDataSetLogicalTableMapDataTransformsProjectOperation `field:"optional" json:"projectOperation" yaml:"projectOperation"`
	// <p>A transform operation that renames a column.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#rename_column_operation QuicksightDataSet#rename_column_operation}
	RenameColumnOperation *QuicksightDataSetLogicalTableMapDataTransformsRenameColumnOperation `field:"optional" json:"renameColumnOperation" yaml:"renameColumnOperation"`
	// <p>A transform operation that tags a column with additional information.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/quicksight_data_set#tag_column_operation QuicksightDataSet#tag_column_operation}
	TagColumnOperation *QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperation `field:"optional" json:"tagColumnOperation" yaml:"tagColumnOperation"`
}

