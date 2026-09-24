// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package quicksightcustomization


type QuicksightCustomizationTags struct {
	// Tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_customization#key QuicksightCustomization#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/quicksight_customization#value QuicksightCustomization#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

