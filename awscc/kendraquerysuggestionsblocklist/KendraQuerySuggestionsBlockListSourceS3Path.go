// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendraquerysuggestionsblocklist


type KendraQuerySuggestionsBlockListSourceS3Path struct {
	// The name of the S3 bucket that contains the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_query_suggestions_block_list#bucket KendraQuerySuggestionsBlockList#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// The name of the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_query_suggestions_block_list#key KendraQuerySuggestionsBlockList#key}
	Key *string `field:"required" json:"key" yaml:"key"`
}

