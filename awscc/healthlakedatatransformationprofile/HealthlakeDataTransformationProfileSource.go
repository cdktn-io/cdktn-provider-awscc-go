// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakedatatransformationprofile


type HealthlakeDataTransformationProfileSource struct {
	// Create the profile by cloning a specific version of an existing profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/healthlake_data_transformation_profile#existing_versioned_profile_id HealthlakeDataTransformationProfile#existing_versioned_profile_id}
	ExistingVersionedProfileId *HealthlakeDataTransformationProfileSourceExistingVersionedProfileId `field:"optional" json:"existingVersionedProfileId" yaml:"existingVersionedProfileId"`
	// Create the profile from raw Velocity template mapping content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/healthlake_data_transformation_profile#profile_mapping HealthlakeDataTransformationProfile#profile_mapping}
	ProfileMapping *HealthlakeDataTransformationProfileSourceProfileMapping `field:"optional" json:"profileMapping" yaml:"profileMapping"`
	// Create the profile from a predefined starter profile of transformation templates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/healthlake_data_transformation_profile#starter_profile HealthlakeDataTransformationProfile#starter_profile}
	StarterProfile *HealthlakeDataTransformationProfileSourceStarterProfile `field:"optional" json:"starterProfile" yaml:"starterProfile"`
}

