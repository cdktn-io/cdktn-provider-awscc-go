// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dynamodbglobaltable


type DynamodbGlobalTableReplicas struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#region DynamodbGlobalTable#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#contributor_insights_specification DynamodbGlobalTable#contributor_insights_specification}.
	ContributorInsightsSpecification *DynamodbGlobalTableReplicasContributorInsightsSpecification `field:"optional" json:"contributorInsightsSpecification" yaml:"contributorInsightsSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#deletion_protection_enabled DynamodbGlobalTable#deletion_protection_enabled}.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#global_secondary_indexes DynamodbGlobalTable#global_secondary_indexes}.
	GlobalSecondaryIndexes interface{} `field:"optional" json:"globalSecondaryIndexes" yaml:"globalSecondaryIndexes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#global_table_settings_replication_mode DynamodbGlobalTable#global_table_settings_replication_mode}.
	GlobalTableSettingsReplicationMode *string `field:"optional" json:"globalTableSettingsReplicationMode" yaml:"globalTableSettingsReplicationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#kinesis_stream_specification DynamodbGlobalTable#kinesis_stream_specification}.
	KinesisStreamSpecification *DynamodbGlobalTableReplicasKinesisStreamSpecification `field:"optional" json:"kinesisStreamSpecification" yaml:"kinesisStreamSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#point_in_time_recovery_specification DynamodbGlobalTable#point_in_time_recovery_specification}.
	PointInTimeRecoverySpecification *DynamodbGlobalTableReplicasPointInTimeRecoverySpecification `field:"optional" json:"pointInTimeRecoverySpecification" yaml:"pointInTimeRecoverySpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#read_on_demand_throughput_settings DynamodbGlobalTable#read_on_demand_throughput_settings}.
	ReadOnDemandThroughputSettings *DynamodbGlobalTableReplicasReadOnDemandThroughputSettings `field:"optional" json:"readOnDemandThroughputSettings" yaml:"readOnDemandThroughputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#read_provisioned_throughput_settings DynamodbGlobalTable#read_provisioned_throughput_settings}.
	ReadProvisionedThroughputSettings *DynamodbGlobalTableReplicasReadProvisionedThroughputSettings `field:"optional" json:"readProvisionedThroughputSettings" yaml:"readProvisionedThroughputSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#replica_stream_specification DynamodbGlobalTable#replica_stream_specification}.
	ReplicaStreamSpecification *DynamodbGlobalTableReplicasReplicaStreamSpecification `field:"optional" json:"replicaStreamSpecification" yaml:"replicaStreamSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#resource_policy DynamodbGlobalTable#resource_policy}.
	ResourcePolicy *DynamodbGlobalTableReplicasResourcePolicy `field:"optional" json:"resourcePolicy" yaml:"resourcePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#sse_specification DynamodbGlobalTable#sse_specification}.
	SseSpecification *DynamodbGlobalTableReplicasSseSpecification `field:"optional" json:"sseSpecification" yaml:"sseSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#table_class DynamodbGlobalTable#table_class}.
	TableClass *string `field:"optional" json:"tableClass" yaml:"tableClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/dynamodb_global_table#tags DynamodbGlobalTable#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

