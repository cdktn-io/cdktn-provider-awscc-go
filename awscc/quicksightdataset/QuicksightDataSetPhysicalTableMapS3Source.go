// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetPhysicalTableMapS3Source struct {
	// <p>The amazon Resource Name (ARN) for the data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_set#data_source_arn QuicksightDataSet#data_source_arn}
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// <p>A physical table type for as S3 data source.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_set#input_columns QuicksightDataSet#input_columns}
	InputColumns interface{} `field:"optional" json:"inputColumns" yaml:"inputColumns"`
	// <p>Information about the format for a source file or files.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/quicksight_data_set#upload_settings QuicksightDataSet#upload_settings}
	UploadSettings *QuicksightDataSetPhysicalTableMapS3SourceUploadSettings `field:"optional" json:"uploadSettings" yaml:"uploadSettings"`
}

