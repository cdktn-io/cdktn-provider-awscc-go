// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableKinesisStreamSpecification struct {
	// The precision for the time and date that the stream was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_table#approximate_creation_date_time_precision DynamodbTable#approximate_creation_date_time_precision}
	ApproximateCreationDateTimePrecision *string `field:"optional" json:"approximateCreationDateTimePrecision" yaml:"approximateCreationDateTimePrecision"`
	// The ARN for a specific Kinesis data stream.  Length Constraints: Minimum length of 37. Maximum length of 1024.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_table#stream_arn DynamodbTable#stream_arn}
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

