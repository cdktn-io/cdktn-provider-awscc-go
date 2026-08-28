// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakedatatransformationprofile


type HealthlakeDataTransformationProfileSourceStarterProfile struct {
	// The name of the starter profile to seed the profile from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/healthlake_data_transformation_profile#starter_profile_name HealthlakeDataTransformationProfile#starter_profile_name}
	StarterProfileName *string `field:"optional" json:"starterProfileName" yaml:"starterProfileName"`
}

