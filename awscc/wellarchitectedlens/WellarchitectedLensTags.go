// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package wellarchitectedlens


type WellarchitectedLensTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wellarchitected_lens#key WellarchitectedLens#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/wellarchitected_lens#value WellarchitectedLens#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

