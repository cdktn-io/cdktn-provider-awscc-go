// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorcrossaccountattachment


type GlobalacceleratorCrossAccountAttachmentTags struct {
	// Key of the tag. Value can be 1 to 127 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/globalaccelerator_cross_account_attachment#key GlobalacceleratorCrossAccountAttachment#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Value for the tag. Value can be 1 to 255 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/globalaccelerator_cross_account_attachment#value GlobalacceleratorCrossAccountAttachment#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

