// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecificationSchedulingConfig struct {
	// Configuration for balancing inference component copies across Availability Zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#availability_zone_balance SagemakerInferenceComponent#availability_zone_balance}
	AvailabilityZoneBalance *SagemakerInferenceComponentSpecificationSchedulingConfigAvailabilityZoneBalance `field:"optional" json:"availabilityZoneBalance" yaml:"availabilityZoneBalance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#placement_strategy SagemakerInferenceComponent#placement_strategy}.
	PlacementStrategy *string `field:"optional" json:"placementStrategy" yaml:"placementStrategy"`
}

