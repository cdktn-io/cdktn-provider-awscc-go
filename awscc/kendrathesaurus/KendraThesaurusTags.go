// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendrathesaurus


type KendraThesaurusTags struct {
	// The key for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_thesaurus#key KendraThesaurus#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value associated with the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/kendra_thesaurus#value KendraThesaurus#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

