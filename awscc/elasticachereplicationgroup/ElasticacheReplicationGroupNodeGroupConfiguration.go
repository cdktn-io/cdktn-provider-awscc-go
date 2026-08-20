// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachereplicationgroup


type ElasticacheReplicationGroupNodeGroupConfiguration struct {
	// Either the ElastiCache for Redis supplied 4-digit id or a user supplied id for the node group these configuration values apply to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_replication_group#node_group_id ElasticacheReplicationGroup#node_group_id}
	NodeGroupId *string `field:"optional" json:"nodeGroupId" yaml:"nodeGroupId"`
	// The Availability Zone where the primary node of this node group (shard) is launched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_replication_group#primary_availability_zone ElasticacheReplicationGroup#primary_availability_zone}
	PrimaryAvailabilityZone *string `field:"optional" json:"primaryAvailabilityZone" yaml:"primaryAvailabilityZone"`
	// A list of Availability Zones to be used for the read replicas.
	//
	// The number of Availability Zones in this list must match the value of ReplicaCount or ReplicasPerNodeGroup if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_replication_group#replica_availability_zones ElasticacheReplicationGroup#replica_availability_zones}
	ReplicaAvailabilityZones *[]*string `field:"optional" json:"replicaAvailabilityZones" yaml:"replicaAvailabilityZones"`
	// The number of read replica nodes in this node group (shard).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_replication_group#replica_count ElasticacheReplicationGroup#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
	// A string of comma-separated values where the first set of values are the slot numbers (zero based), and the second set of values are the keyspaces for each slot.
	//
	// The following example specifies three slots (numbered 0, 1, and 2): 0,1,2,0-4999,5000-9999,10000-16,383.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/elasticache_replication_group#slots ElasticacheReplicationGroup#slots}
	Slots *string `field:"optional" json:"slots" yaml:"slots"`
}

