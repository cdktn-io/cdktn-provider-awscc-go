// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package memorydbparametergroup


type MemorydbParameterGroupTags struct {
	// The key for the tag. May not be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/memorydb_parameter_group#key MemorydbParameterGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value. May be null.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/memorydb_parameter_group#value MemorydbParameterGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

