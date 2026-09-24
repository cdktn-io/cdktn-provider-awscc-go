// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakedatatransformationprofile


type HealthlakeDataTransformationProfileSourceExistingVersionedProfileId struct {
	// The unique identifier of the source profile to clone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/healthlake_data_transformation_profile#profile_id HealthlakeDataTransformationProfile#profile_id}
	ProfileId *string `field:"optional" json:"profileId" yaml:"profileId"`
	// The version number of the source profile to clone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/healthlake_data_transformation_profile#version HealthlakeDataTransformationProfile#version}
	Version *float64 `field:"optional" json:"version" yaml:"version"`
}

