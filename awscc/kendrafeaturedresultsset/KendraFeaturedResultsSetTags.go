// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendrafeaturedresultsset


type KendraFeaturedResultsSetTags struct {
	// The key for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_featured_results_set#key KendraFeaturedResultsSet#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value associated with the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kendra_featured_results_set#value KendraFeaturedResultsSet#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

