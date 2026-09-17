// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cassandratable


type CassandraTableClusteringKeyColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cassandra_table#column CassandraTable#column}.
	Column *CassandraTableClusteringKeyColumnsColumn `field:"optional" json:"column" yaml:"column"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/cassandra_table#order_by CassandraTable#order_by}.
	OrderBy *string `field:"optional" json:"orderBy" yaml:"orderBy"`
}

