// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbtable


type DynamodbTableGlobalSecondaryIndexes struct {
	// The settings used to specify whether to enable CloudWatch Contributor Insights for the global table and define which events to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#contributor_insights_specification DynamodbTable#contributor_insights_specification}
	ContributorInsightsSpecification *DynamodbTableGlobalSecondaryIndexesContributorInsightsSpecification `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// The name of the global secondary index. The name must be unique among all other indexes on this table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#index_name DynamodbTable#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:   +  ``HASH`` - partition key   +  ``RANGE`` - sort key      The partition key of an item is also known as its *hash attribute*.
	//
	// The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.
	//  The sort key of an item is also known as its *range attribute*. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#key_schema DynamodbTable#key_schema}
	KeySchema interface{} `field:"optional" json:"keySchema" yaml:"keySchema"`
	// The maximum number of read and write units for the specified global secondary index.
	//
	// If you use this parameter, you must specify ``MaxReadRequestUnits``, ``MaxWriteRequestUnits``, or both. You must use either ``OnDemandThroughput`` or ``ProvisionedThroughput`` based on your table's capacity mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#on_demand_throughput DynamodbTable#on_demand_throughput}
	OnDemandThroughput *DynamodbTableGlobalSecondaryIndexesOnDemandThroughput `field:"optional" json:"onDemandThroughput" yaml:"onDemandThroughput"`
	// Represents attributes that are copied (projected) from the table into the global secondary index.
	//
	// These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#projection DynamodbTable#projection}
	Projection *DynamodbTableGlobalSecondaryIndexesProjection `field:"optional" json:"projection" yaml:"projection"`
	// Represents the provisioned throughput settings for the specified global secondary index.
	//
	// You must use either ``OnDemandThroughput`` or ``ProvisionedThroughput`` based on your table's capacity mode.
	//  For current minimum and maximum provisioned throughput values, see [Service, Account, and Table Quotas](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html) in the *Amazon DynamoDB Developer Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#provisioned_throughput DynamodbTable#provisioned_throughput}
	ProvisionedThroughput *DynamodbTableGlobalSecondaryIndexesProvisionedThroughput `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
	// Represents the warm throughput value (in read units per second and write units per second) for the specified secondary index.
	//
	// If you use this parameter, you must specify ``ReadUnitsPerSecond``, ``WriteUnitsPerSecond``, or both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/dynamodb_table#warm_throughput DynamodbTable#warm_throughput}
	WarmThroughput *DynamodbTableGlobalSecondaryIndexesWarmThroughput `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
}

