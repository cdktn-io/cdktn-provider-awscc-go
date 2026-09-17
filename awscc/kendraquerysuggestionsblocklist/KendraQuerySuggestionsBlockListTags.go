// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraquerysuggestionsblocklist


type KendraQuerySuggestionsBlockListTags struct {
	// The key for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_query_suggestions_block_list#key KendraQuerySuggestionsBlockList#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value associated with the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_query_suggestions_block_list#value KendraQuerySuggestionsBlockList#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

