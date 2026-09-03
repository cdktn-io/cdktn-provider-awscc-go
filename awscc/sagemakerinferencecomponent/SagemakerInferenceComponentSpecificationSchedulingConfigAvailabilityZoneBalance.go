// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecificationSchedulingConfigAvailabilityZoneBalance struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_inference_component#enforcement_mode SagemakerInferenceComponent#enforcement_mode}.
	EnforcementMode *string `field:"optional" json:"enforcementMode" yaml:"enforcementMode"`
	// The maximum allowed difference in the number of inference component copies between any two Availability Zones.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_inference_component#max_imbalance SagemakerInferenceComponent#max_imbalance}
	MaxImbalance *float64 `field:"optional" json:"maxImbalance" yaml:"maxImbalance"`
}

