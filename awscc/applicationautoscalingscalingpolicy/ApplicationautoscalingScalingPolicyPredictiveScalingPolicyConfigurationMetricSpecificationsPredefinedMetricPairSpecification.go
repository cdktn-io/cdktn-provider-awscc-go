// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedMetricPairSpecification struct {
	// Indicates which metrics to use.
	//
	// There are two different types of metrics for each metric type: one is a load metric and one is a scaling metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationautoscaling_scaling_policy#predefined_metric_type ApplicationautoscalingScalingPolicy#predefined_metric_type}
	PredefinedMetricType *string `field:"optional" json:"predefinedMetricType" yaml:"predefinedMetricType"`
	// A label that uniquely identifies a specific target group from which to determine the total and average request count.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationautoscaling_scaling_policy#resource_label ApplicationautoscalingScalingPolicy#resource_label}
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

