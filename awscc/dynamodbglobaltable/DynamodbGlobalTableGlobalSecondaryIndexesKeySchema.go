// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableGlobalSecondaryIndexesKeySchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_global_table#attribute_name DynamodbGlobalTable#attribute_name}.
	AttributeName *string `field:"optional" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/dynamodb_global_table#key_type DynamodbGlobalTable#key_type}.
	KeyType *string `field:"optional" json:"keyType" yaml:"keyType"`
}

