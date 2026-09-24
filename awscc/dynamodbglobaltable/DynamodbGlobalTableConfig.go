// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DynamodbGlobalTableConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#replicas DynamodbGlobalTable#replicas}.
	Replicas interface{} `field:"required" json:"replicas" yaml:"replicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#attribute_definitions DynamodbGlobalTable#attribute_definitions}.
	AttributeDefinitions interface{} `field:"optional" json:"attributeDefinitions" yaml:"attributeDefinitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#billing_mode DynamodbGlobalTable#billing_mode}.
	BillingMode *string `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#global_secondary_indexes DynamodbGlobalTable#global_secondary_indexes}.
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#global_table_source_arn DynamodbGlobalTable#global_table_source_arn}.
	GlobalTableSourceArn *string `field:"optional" json:"globalTableSourceArn" yaml:"globalTableSourceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#global_table_witnesses DynamodbGlobalTable#global_table_witnesses}.
	GlobalTableWitnesses interface{} `field:"optional" json:"globalTableWitnesses" yaml:"globalTableWitnesses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#key_schema DynamodbGlobalTable#key_schema}.
	KeySchema interface{} `field:"optional" json:"keySchema" yaml:"keySchema"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#local_secondary_indexes DynamodbGlobalTable#local_secondary_indexes}.
	LocalSecondaryIndexes interface{} `field:"optional" json:"localSecondaryIndexes" yaml:"localSecondaryIndexes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#multi_region_consistency DynamodbGlobalTable#multi_region_consistency}.
	MultiRegionConsistency *string `field:"optional" json:"multiRegionConsistency" yaml:"multiRegionConsistency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#read_on_demand_throughput_settings DynamodbGlobalTable#read_on_demand_throughput_settings}.
	ReadOnDemandThroughputSettings *DynamodbGlobalTableReadOnDemandThroughputSettings `field:"optional" json:"readOnDemandThroughputSettings" yaml:"readOnDemandThroughputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#read_provisioned_throughput_settings DynamodbGlobalTable#read_provisioned_throughput_settings}.
	ReadProvisionedThroughputSettings *DynamodbGlobalTableReadProvisionedThroughputSettings `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#sse_specification DynamodbGlobalTable#sse_specification}.
	SseSpecification *DynamodbGlobalTableSseSpecification `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#stream_specification DynamodbGlobalTable#stream_specification}.
	StreamSpecification *DynamodbGlobalTableStreamSpecification `field:"optional" json:"streamSpecification" yaml:"streamSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#table_name DynamodbGlobalTable#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#time_to_live_specification DynamodbGlobalTable#time_to_live_specification}.
	TimeToLiveSpecification *DynamodbGlobalTableTimeToLiveSpecification `field:"optional" json:"timeToLiveSpecification" yaml:"timeToLiveSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#vector_indexes DynamodbGlobalTable#vector_indexes}.
	VectorIndexes interface{} `field:"optional" json:"vectorIndexes" yaml:"vectorIndexes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#warm_throughput DynamodbGlobalTable#warm_throughput}.
	WarmThroughput *DynamodbGlobalTableWarmThroughput `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#write_on_demand_throughput_settings DynamodbGlobalTable#write_on_demand_throughput_settings}.
	WriteOnDemandThroughputSettings *DynamodbGlobalTableWriteOnDemandThroughputSettings `field:"optional" json:"writeOnDemandThroughputSettings" yaml:"writeOnDemandThroughputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/dynamodb_global_table#write_provisioned_throughput_settings DynamodbGlobalTable#write_provisioned_throughput_settings}.
	WriteProvisionedThroughputSettings *DynamodbGlobalTableWriteProvisionedThroughputSettings `field:"optional" json:"writeProvisionedThroughputSettings" yaml:"writeProvisionedThroughputSettings"`
}

