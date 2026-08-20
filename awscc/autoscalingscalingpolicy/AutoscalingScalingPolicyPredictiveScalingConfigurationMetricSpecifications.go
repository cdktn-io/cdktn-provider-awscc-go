// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingscalingpolicy


type AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_scaling_policy#customized_capacity_metric_specification AutoscalingScalingPolicy#customized_capacity_metric_specification}.
	CustomizedCapacityMetricSpecification *AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsCustomizedCapacityMetricSpecification `field:"optional" json:"customizedCapacityMetricSpecification" yaml:"customizedCapacityMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_scaling_policy#customized_load_metric_specification AutoscalingScalingPolicy#customized_load_metric_specification}.
	CustomizedLoadMetricSpecification *AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsCustomizedLoadMetricSpecification `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_scaling_policy#customized_scaling_metric_specification AutoscalingScalingPolicy#customized_scaling_metric_specification}.
	CustomizedScalingMetricSpecification *AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsCustomizedScalingMetricSpecification `field:"optional" json:"customizedScalingMetricSpecification" yaml:"customizedScalingMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_scaling_policy#predefined_load_metric_specification AutoscalingScalingPolicy#predefined_load_metric_specification}.
	PredefinedLoadMetricSpecification *AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsPredefinedLoadMetricSpecification `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_scaling_policy#predefined_metric_pair_specification AutoscalingScalingPolicy#predefined_metric_pair_specification}.
	PredefinedMetricPairSpecification *AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsPredefinedMetricPairSpecification `field:"optional" json:"predefinedMetricPairSpecification" yaml:"predefinedMetricPairSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_scaling_policy#predefined_scaling_metric_specification AutoscalingScalingPolicy#predefined_scaling_metric_specification}.
	PredefinedScalingMetricSpecification *AutoscalingScalingPolicyPredictiveScalingConfigurationMetricSpecificationsPredefinedScalingMetricSpecification `field:"optional" json:"predefinedScalingMetricSpecification" yaml:"predefinedScalingMetricSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/autoscaling_scaling_policy#target_value AutoscalingScalingPolicy#target_value}.
	TargetValue *float64 `field:"optional" json:"targetValue" yaml:"targetValue"`
}

