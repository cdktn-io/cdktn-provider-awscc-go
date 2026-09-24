// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CassandraTableConfig struct {
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
	// Name for Cassandra keyspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#keyspace_name CassandraTable#keyspace_name}
	KeyspaceName *string `field:"required" json:"keyspaceName" yaml:"keyspaceName"`
	// Partition key columns of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#partition_key_columns CassandraTable#partition_key_columns}
	PartitionKeyColumns interface{} `field:"required" json:"partitionKeyColumns" yaml:"partitionKeyColumns"`
	// Represents the read and write settings used for AutoScaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#auto_scaling_specifications CassandraTable#auto_scaling_specifications}
	AutoScalingSpecifications *CassandraTableAutoScalingSpecifications `field:"optional" json:"autoScalingSpecifications" yaml:"autoScalingSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#billing_mode CassandraTable#billing_mode}.
	BillingMode *CassandraTableBillingMode `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Represents the CDC configuration for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#cdc_specification CassandraTable#cdc_specification}
	CdcSpecification *CassandraTableCdcSpecification `field:"optional" json:"cdcSpecification" yaml:"cdcSpecification"`
	// Indicates whether client side timestamps are enabled (true) or disabled (false) on the table.
	//
	// False by default, once it is enabled it cannot be disabled again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#client_side_timestamps_enabled CassandraTable#client_side_timestamps_enabled}
	ClientSideTimestampsEnabled interface{} `field:"optional" json:"clientSideTimestampsEnabled" yaml:"clientSideTimestampsEnabled"`
	// Clustering key columns of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#clustering_key_columns CassandraTable#clustering_key_columns}
	ClusteringKeyColumns interface{} `field:"optional" json:"clusteringKeyColumns" yaml:"clusteringKeyColumns"`
	// Default TTL (Time To Live) in seconds, where zero is disabled.
	//
	// If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#default_time_to_live CassandraTable#default_time_to_live}
	DefaultTimeToLive *float64 `field:"optional" json:"defaultTimeToLive" yaml:"defaultTimeToLive"`
	// Represents the settings used to enable server-side encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#encryption_specification CassandraTable#encryption_specification}
	EncryptionSpecification *CassandraTableEncryptionSpecification `field:"optional" json:"encryptionSpecification" yaml:"encryptionSpecification"`
	// Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#point_in_time_recovery_enabled CassandraTable#point_in_time_recovery_enabled}
	PointInTimeRecoveryEnabled interface{} `field:"optional" json:"pointInTimeRecoveryEnabled" yaml:"pointInTimeRecoveryEnabled"`
	// Non-key columns of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#regular_columns CassandraTable#regular_columns}
	RegularColumns interface{} `field:"optional" json:"regularColumns" yaml:"regularColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#replica_specifications CassandraTable#replica_specifications}.
	ReplicaSpecifications interface{} `field:"optional" json:"replicaSpecifications" yaml:"replicaSpecifications"`
	// Name for Cassandra table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#table_name CassandraTable#table_name}
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#tags CassandraTable#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Warm throughput configuration for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#warm_throughput CassandraTable#warm_throughput}
	WarmThroughput *CassandraTableWarmThroughput `field:"optional" json:"warmThroughput" yaml:"warmThroughput"`
}

