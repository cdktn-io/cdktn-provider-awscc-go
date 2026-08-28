// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable


type CassandraTableClusteringKeyColumnsColumn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cassandra_table#column_name CassandraTable#column_name}.
	ColumnName *string `field:"optional" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cassandra_table#column_type CassandraTable#column_type}.
	ColumnType *string `field:"optional" json:"columnType" yaml:"columnType"`
}

