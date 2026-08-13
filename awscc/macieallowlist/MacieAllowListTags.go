// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package macieallowlist


type MacieAllowListTags struct {
	// The tag's key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/macie_allow_list#key MacieAllowList#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag's value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/macie_allow_list#value MacieAllowList#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

