// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package gluepartition


type GluePartitionPartitionInput struct {
	// The values of the partition.
	//
	// Although this parameter is not required by the SDK, you must specify this parameter for a valid input. The values for the keys for the new partition must be passed as an array of String objects that must be ordered in the same order as the partition keys appearing in the Amazon S3 prefix. Otherwise AWS Glue will add the values to the wrong keys.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#values GluePartition#values}
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Key-value pairs defining partition parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#parameters GluePartition#parameters}
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
	// Provides information about the physical location where the partition is stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/glue_partition#storage_descriptor GluePartition#storage_descriptor}
	StorageDescriptor *GluePartitionPartitionInputStorageDescriptor `field:"optional" json:"storageDescriptor" yaml:"storageDescriptor"`
}

