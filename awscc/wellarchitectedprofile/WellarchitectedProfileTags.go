// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedprofile


type WellarchitectedProfileTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wellarchitected_profile#key WellarchitectedProfile#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/wellarchitected_profile#value WellarchitectedProfile#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

