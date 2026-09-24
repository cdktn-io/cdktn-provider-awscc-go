// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable


type CassandraTableReplicaSpecifications struct {
	// Represents configuration for auto scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#read_capacity_auto_scaling CassandraTable#read_capacity_auto_scaling}
	ReadCapacityAutoScaling *CassandraTableReplicaSpecificationsReadCapacityAutoScaling `field:"optional" json:"readCapacityAutoScaling" yaml:"readCapacityAutoScaling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#read_capacity_units CassandraTable#read_capacity_units}.
	ReadCapacityUnits *float64 `field:"optional" json:"readCapacityUnits" yaml:"readCapacityUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#region CassandraTable#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

