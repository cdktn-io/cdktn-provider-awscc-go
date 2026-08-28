// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheglobalreplicationgroup


type ElasticacheGlobalReplicationGroupMembers struct {
	// Regionally unique identifier for the member i.e. ReplicationGroupId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_global_replication_group#replication_group_id ElasticacheGlobalReplicationGroup#replication_group_id}
	ReplicationGroupId *string `field:"optional" json:"replicationGroupId" yaml:"replicationGroupId"`
	// The AWS region of the Global Datastore member.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_global_replication_group#replication_group_region ElasticacheGlobalReplicationGroup#replication_group_region}
	ReplicationGroupRegion *string `field:"optional" json:"replicationGroupRegion" yaml:"replicationGroupRegion"`
	// Indicates the role of the member, primary or secondary.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_global_replication_group#role ElasticacheGlobalReplicationGroup#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

