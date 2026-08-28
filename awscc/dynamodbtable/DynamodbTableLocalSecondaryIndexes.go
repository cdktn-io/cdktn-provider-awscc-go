// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableLocalSecondaryIndexes struct {
	// The name of the local secondary index. The name must be unique among all other indexes on this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#index_name DynamodbTable#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The complete key schema for the local secondary index, consisting of one or more pairs of attribute names and key types:   +  ``HASH`` - partition key   +  ``RANGE`` - sort key      The partition key of an item is also known as its *hash attribute*.
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	//  The sort key of an item is also known as its *range attribute*. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#key_schema DynamodbTable#key_schema}
	KeySchema interface{} `field:"optional" json:"keySchema" yaml:"keySchema"`
	// Represents attributes that are copied (projected) from the table into the local secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/dynamodb_table#projection DynamodbTable#projection}
	Projection *DynamodbTableLocalSecondaryIndexesProjection `field:"optional" json:"projection" yaml:"projection"`
}

