// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ssmdocument


type SsmDocumentTags struct {
	// The name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssm_document#key SsmDocument#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ssm_document#value SsmDocument#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

