// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightdatasource


type QuicksightDataSourceDataSourceParametersS3ParametersManifestFileLocation struct {
	// <p>Amazon S3 bucket.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_data_source#bucket QuicksightDataSource#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// <p>Amazon S3 key that identifies an object.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/quicksight_data_source#key QuicksightDataSource#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
}

