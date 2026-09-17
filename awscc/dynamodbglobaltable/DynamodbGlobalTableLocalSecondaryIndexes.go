// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableLocalSecondaryIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#index_name DynamodbGlobalTable#index_name}.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#key_schema DynamodbGlobalTable#key_schema}.
	KeySchema interface{} `field:"optional" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_global_table#projection DynamodbGlobalTable#projection}.
	Projection *DynamodbGlobalTableLocalSecondaryIndexesProjection `field:"optional" json:"projection" yaml:"projection"`
}

