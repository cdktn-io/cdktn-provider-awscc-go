// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableReplicasKinesisStreamSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#approximate_creation_date_time_precision DynamodbGlobalTable#approximate_creation_date_time_precision}.
	ApproximateCreationDateTimePrecision *string `field:"optional" json:"approximateCreationDateTimePrecision" yaml:"approximateCreationDateTimePrecision"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#stream_arn DynamodbGlobalTable#stream_arn}.
	StreamArn *string `field:"optional" json:"streamArn" yaml:"streamArn"`
}

