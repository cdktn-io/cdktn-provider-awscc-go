// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceDataSourceConfigurationS3Configuration struct {
	// The ARN of the bucket that contains the data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#bucket_arn BedrockDataSource#bucket_arn}
	BucketArn *string `field:"optional" json:"bucketArn" yaml:"bucketArn"`
	// The account ID for the owner of the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#bucket_owner_account_id BedrockDataSource#bucket_owner_account_id}
	BucketOwnerAccountId *string `field:"optional" json:"bucketOwnerAccountId" yaml:"bucketOwnerAccountId"`
	// A list of S3 prefixes that define the object containing the data sources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#inclusion_prefixes BedrockDataSource#inclusion_prefixes}
	InclusionPrefixes *[]*string `field:"optional" json:"inclusionPrefixes" yaml:"inclusionPrefixes"`
}

