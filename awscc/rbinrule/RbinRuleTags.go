// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rbinrule


type RbinRuleTags struct {
	// A unique identifier for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#key RbinRule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// String which you can use to describe or define the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/rbin_rule#value RbinRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

