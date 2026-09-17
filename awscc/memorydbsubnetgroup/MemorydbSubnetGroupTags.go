// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package memorydbsubnetgroup


type MemorydbSubnetGroupTags struct {
	// The key for the tag. May not be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/memorydb_subnet_group#key MemorydbSubnetGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value. May be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/memorydb_subnet_group#value MemorydbSubnetGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

