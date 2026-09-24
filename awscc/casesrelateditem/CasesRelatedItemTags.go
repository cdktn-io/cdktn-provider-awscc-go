// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casesrelateditem


type CasesRelatedItemTags struct {
	// The key of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cases_related_item#key CasesRelatedItem#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cases_related_item#value CasesRelatedItem#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

