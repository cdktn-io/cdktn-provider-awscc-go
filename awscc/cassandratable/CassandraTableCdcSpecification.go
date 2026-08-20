// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable


type CassandraTableCdcSpecification struct {
	// Indicates whether CDC is enabled or disabled for the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cassandra_table#status CassandraTable#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// An array of key-value pairs to apply to the CDC stream resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cassandra_table#tags CassandraTable#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// Specifies what data should be captured in the change data stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cassandra_table#view_type CassandraTable#view_type}
	ViewType *string `field:"optional" json:"viewType" yaml:"viewType"`
}

