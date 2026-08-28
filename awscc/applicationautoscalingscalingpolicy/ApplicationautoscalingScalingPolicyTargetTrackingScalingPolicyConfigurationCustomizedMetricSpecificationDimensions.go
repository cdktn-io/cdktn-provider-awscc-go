// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimensions struct {
	// The name of the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationautoscaling_scaling_policy#name ApplicationautoscalingScalingPolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationautoscaling_scaling_policy#value ApplicationautoscalingScalingPolicy#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

