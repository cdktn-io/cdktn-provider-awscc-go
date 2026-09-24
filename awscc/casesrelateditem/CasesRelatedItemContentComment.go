// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package casesrelateditem


type CasesRelatedItemContentComment struct {
	// Text in the body of a comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cases_related_item#body CasesRelatedItem#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Type of the text in the comment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/cases_related_item#content_type CasesRelatedItem#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

