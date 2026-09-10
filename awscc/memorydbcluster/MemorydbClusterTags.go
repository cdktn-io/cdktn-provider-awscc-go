// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package memorydbcluster


type MemorydbClusterTags struct {
	// The key for the tag. May not be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#key MemorydbCluster#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value. May be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/memorydb_cluster#value MemorydbCluster#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

