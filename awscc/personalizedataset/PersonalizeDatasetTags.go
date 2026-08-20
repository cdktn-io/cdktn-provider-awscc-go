// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package personalizedataset


type PersonalizeDatasetTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_dataset#key PersonalizeDataset#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/personalize_dataset#value PersonalizeDataset#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

