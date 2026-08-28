// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackset


type CloudformationStackSetTags struct {
	// A string used to identify this tag. You can specify a maximum of 127 characters for a tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#key CloudformationStackSet#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A string containing the value for this tag.
	//
	// You can specify a maximum of 256 characters for a tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/cloudformation_stack_set#value CloudformationStackSet#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

