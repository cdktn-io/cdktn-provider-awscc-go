// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandrakeyspace

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CassandraKeyspaceConfig struct {
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
	// Indicates whether client-side timestamps are enabled (true) or disabled (false) for all tables in the keyspace.
	//
	// To add a Region to a single-Region keyspace with at least one table, the value must be set to true. After you enabled client-side timestamps for a table, you can?t disable it again.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cassandra_keyspace#client_side_timestamps_enabled CassandraKeyspace#client_side_timestamps_enabled}
	ClientSideTimestampsEnabled interface{} `field:"optional" json:"clientSideTimestampsEnabled" yaml:"clientSideTimestampsEnabled"`
	// Name for Cassandra keyspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cassandra_keyspace#keyspace_name CassandraKeyspace#keyspace_name}
	KeyspaceName *string `field:"optional" json:"keyspaceName" yaml:"keyspaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cassandra_keyspace#replication_specification CassandraKeyspace#replication_specification}.
	ReplicationSpecification *CassandraKeyspaceReplicationSpecification `field:"optional" json:"replicationSpecification" yaml:"replicationSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cassandra_keyspace#tags CassandraKeyspace#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

