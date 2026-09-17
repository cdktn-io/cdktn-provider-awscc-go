// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecifications struct {
	// The customized capacity metric specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#customized_capacity_metric_specification ApplicationautoscalingScalingPolicy#customized_capacity_metric_specification}
	CustomizedCapacityMetricSpecification *ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsCustomizedCapacityMetricSpecification `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// The customized load metric specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#customized_load_metric_specification ApplicationautoscalingScalingPolicy#customized_load_metric_specification}
	CustomizedLoadMetricSpecification *ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsCustomizedLoadMetricSpecification `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// The customized scaling metric specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#customized_scaling_metric_specification ApplicationautoscalingScalingPolicy#customized_scaling_metric_specification}
	CustomizedScalingMetricSpecification *ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsCustomizedScalingMetricSpecification `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// The predefined load metric specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#predefined_load_metric_specification ApplicationautoscalingScalingPolicy#predefined_load_metric_specification}
	PredefinedLoadMetricSpecification *ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedLoadMetricSpecification `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// The predefined metric pair specification that determines the appropriate scaling metric and load metric to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#predefined_metric_pair_specification ApplicationautoscalingScalingPolicy#predefined_metric_pair_specification}
	PredefinedMetricPairSpecification *ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedMetricPairSpecification `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// The predefined scaling metric specification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#predefined_scaling_metric_specification ApplicationautoscalingScalingPolicy#predefined_scaling_metric_specification}
	PredefinedScalingMetricSpecification *ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfigurationMetricSpecificationsPredefinedScalingMetricSpecification `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
	// Specifies the target utilization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#target_value ApplicationautoscalingScalingPolicy#target_value}
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

