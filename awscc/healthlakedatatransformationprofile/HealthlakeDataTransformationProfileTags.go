// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakedatatransformationprofile


type HealthlakeDataTransformationProfileTags struct {
	// The key of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_data_transformation_profile#key HealthlakeDataTransformationProfile#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_data_transformation_profile#value HealthlakeDataTransformationProfile#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

