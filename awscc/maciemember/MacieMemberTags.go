// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package maciemember


type MacieMemberTags struct {
	// The key of the tag. The maximum length of a tag key is 128 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/macie_member#key MacieMember#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag. The maximum length of a tag value is 256 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/macie_member#value MacieMember#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

