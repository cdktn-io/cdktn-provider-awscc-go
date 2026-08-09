// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sestenant


type SesTenantTags struct {
	// The key of the key-value tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_tenant#key SesTenant#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the key-value tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/ses_tenant#value SesTenant#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

