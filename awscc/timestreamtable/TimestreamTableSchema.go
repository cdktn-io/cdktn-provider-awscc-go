// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package timestreamtable


type TimestreamTableSchema struct {
	// A list of partition keys defining the attributes used to partition the table data.
	//
	// The order of the list determines the partition hierarchy. The name and type of each partition key as well as the partition key order cannot be changed after the table is created. However, the enforcement level of each partition key can be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/timestream_table#composite_partition_key TimestreamTable#composite_partition_key}
	CompositePartitionKey interface{} `field:"optional" json:"compositePartitionKey" yaml:"compositePartitionKey"`
}

