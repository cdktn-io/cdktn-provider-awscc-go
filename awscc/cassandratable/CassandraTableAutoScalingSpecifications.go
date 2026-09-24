// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable


type CassandraTableAutoScalingSpecifications struct {
	// Represents configuration for auto scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#read_capacity_auto_scaling CassandraTable#read_capacity_auto_scaling}
	ReadCapacityAutoScaling *CassandraTableAutoScalingSpecificationsReadCapacityAutoScaling `field:"optional" json:"readCapacityAutoScaling" yaml:"readCapacityAutoScaling"`
	// Represents configuration for auto scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cassandra_table#write_capacity_auto_scaling CassandraTable#write_capacity_auto_scaling}
	WriteCapacityAutoScaling *CassandraTableAutoScalingSpecificationsWriteCapacityAutoScaling `field:"optional" json:"writeCapacityAutoScaling" yaml:"writeCapacityAutoScaling"`
}

