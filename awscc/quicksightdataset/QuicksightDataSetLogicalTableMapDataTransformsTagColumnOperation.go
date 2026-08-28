// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetLogicalTableMapDataTransformsTagColumnOperation struct {
	// <p>The column that this operation acts on.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_set#column_name QuicksightDataSet#column_name}
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// <p>The dataset column tag, currently only used for geospatial type tagging.
	//
	// .</p>
	//         <note>
	//             <p>This is not tags for the AWS tagging feature. .</p>
	//         </note>
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/quicksight_data_set#tags QuicksightDataSet#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

