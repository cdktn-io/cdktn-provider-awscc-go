// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package textractadapter


type TextractAdapterTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/textract_adapter#key TextractAdapter#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/textract_adapter#value TextractAdapter#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

