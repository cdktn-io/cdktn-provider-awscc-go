// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package healthlakedatatransformationprofile


type HealthlakeDataTransformationProfileSourceProfileMapping struct {
	// Map of template file paths to their Velocity template content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/healthlake_data_transformation_profile#profile_mapping HealthlakeDataTransformationProfile#profile_mapping}
	ProfileMapping *map[string]*string `field:"optional" json:"profileMapping" yaml:"profileMapping"`
}

