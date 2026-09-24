// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluepartition


type GluePartitionPartitionInputStorageDescriptorSortColumns struct {
	// The name of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#column GluePartition#column}
	Column *string `field:"optional" json:"column" yaml:"column"`
	// Indicates that the column is sorted in ascending order (== 1), or in descending order (==0).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#sort_order GluePartition#sort_order}
	SortOrder *float64 `field:"optional" json:"sortOrder" yaml:"sortOrder"`
}

