// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableSseSpecification struct {
	// The KMS key that should be used for the KMS encryption.
	//
	// To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB key ``alias/aws/dynamodb``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_table#kms_master_key_id DynamodbTable#kms_master_key_id}
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// Indicates whether server-side encryption is done using an AWS managed key or an AWS owned key.
	//
	// If enabled (true), server-side encryption type is set to ``KMS`` and an AWS managed key is used (KMS charges apply). If disabled (false) or not specified, server-side encryption is set to AWS owned key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_table#sse_enabled DynamodbTable#sse_enabled}
	SseEnabled interface{} `field:"optional" json:"sseEnabled" yaml:"sseEnabled"`
	// Server-side encryption type.
	//
	// The only supported value is:
	//   +  ``KMS`` - Server-side encryption that uses KMSlong. The key is stored in your account and is managed by KMS (KMS charges apply).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_table#sse_type DynamodbTable#sse_type}
	SseType *string `field:"optional" json:"sseType" yaml:"sseType"`
}

