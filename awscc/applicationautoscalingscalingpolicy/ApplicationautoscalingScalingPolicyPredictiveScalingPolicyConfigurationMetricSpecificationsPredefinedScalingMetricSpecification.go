// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecification struct {
	// The metric type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationautoscaling_scaling_policy#predefined_metric_type ApplicationautoscalingScalingPolicy#predefined_metric_type}
	PredefinedMetricType *string `field:"optional" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// A label that uniquely identifies a specific target group from which to determine the average request count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationautoscaling_scaling_policy#resource_label ApplicationautoscalingScalingPolicy#resource_label}
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

