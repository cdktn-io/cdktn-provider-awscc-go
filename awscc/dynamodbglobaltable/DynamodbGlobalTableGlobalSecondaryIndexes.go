// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableGlobalSecondaryIndexes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#index_name DynamodbGlobalTable#index_name}.
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#key_schema DynamodbGlobalTable#key_schema}.
	KeySchema interface{} `field:"optional" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#projection DynamodbGlobalTable#projection}.
	Projection *DynamodbGlobalTableGlobalSecondaryIndexesProjection `field:"optional" json:"projection" yaml:"projection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#read_on_demand_throughput_settings DynamodbGlobalTable#read_on_demand_throughput_settings}.
	ReadOnDemandThroughputSettings *DynamodbGlobalTableGlobalSecondaryIndexesReadOnDemandThroughputSettings `field:"optional" json:"readOnDemandThroughputSettings" yaml:"readOnDemandThroughputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#read_provisioned_throughput_settings DynamodbGlobalTable#read_provisioned_throughput_settings}.
	ReadProvisionedThroughputSettings *DynamodbGlobalTableGlobalSecondaryIndexesReadProvisionedThroughputSettings `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#warm_throughput DynamodbGlobalTable#warm_throughput}.
	WarmThroughput *DynamodbGlobalTableGlobalSecondaryIndexesWarmThroughput `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#write_on_demand_throughput_settings DynamodbGlobalTable#write_on_demand_throughput_settings}.
	WriteOnDemandThroughputSettings *DynamodbGlobalTableGlobalSecondaryIndexesWriteOnDemandThroughputSettings `field:"optional" json:"writeOnDemandThroughputSettings" yaml:"writeOnDemandThroughputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/dynamodb_global_table#write_provisioned_throughput_settings DynamodbGlobalTable#write_provisioned_throughput_settings}.
	WriteProvisionedThroughputSettings *DynamodbGlobalTableGlobalSecondaryIndexesWriteProvisionedThroughputSettings `field:"optional" json:"writeProvisionedThroughputSettings" yaml:"writeProvisionedThroughputSettings"`
}

