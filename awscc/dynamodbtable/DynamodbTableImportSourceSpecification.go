// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableImportSourceSpecification struct {
	// Type of compression to be used on the input coming from the imported table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dynamodb_table#input_compression_type DynamodbTable#input_compression_type}
	InputCompressionType *string `field:"optional" json:"inputCompressionType" yaml:"inputCompressionType"`
	// The format of the source data. Valid values for ``ImportFormat`` are ``CSV``, ``DYNAMODB_JSON`` or ``ION``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dynamodb_table#input_format DynamodbTable#input_format}
	InputFormat *string `field:"optional" json:"inputFormat" yaml:"inputFormat"`
	// Additional properties that specify how the input is formatted,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dynamodb_table#input_format_options DynamodbTable#input_format_options}
	InputFormatOptions *DynamodbTableImportSourceSpecificationInputFormatOptions `field:"optional" json:"inputFormatOptions" yaml:"inputFormatOptions"`
	// The S3 bucket that provides the source for the import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/dynamodb_table#s3_bucket_source DynamodbTable#s3_bucket_source}
	S3BucketSource *DynamodbTableImportSourceSpecificationS3BucketSource `field:"optional" json:"s3BucketSource" yaml:"s3BucketSource"`
}

