// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorage struct {
	// An Amazon S3 location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_data_source#s3_location BedrockDataSource#s3_location}
	S3Location *BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorageS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

