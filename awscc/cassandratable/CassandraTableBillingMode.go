// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable


type CassandraTableBillingMode struct {
	// Capacity mode for the specified table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cassandra_table#mode CassandraTable#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cassandra_table#provisioned_throughput CassandraTable#provisioned_throughput}
	ProvisionedThroughput *CassandraTableBillingModeProvisionedThroughput `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

