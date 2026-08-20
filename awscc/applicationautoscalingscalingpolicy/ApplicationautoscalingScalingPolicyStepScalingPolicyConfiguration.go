// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyStepScalingPolicyConfiguration struct {
	// Specifies whether the ``ScalingAdjustment`` value in the ``StepAdjustment`` property is an absolute number or a percentage of the current capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationautoscaling_scaling_policy#adjustment_type ApplicationautoscalingScalingPolicy#adjustment_type}
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// The amount of time, in seconds, to wait for a previous scaling activity to take effect.
	//
	// If not specified, the default value is 300. For more information, see [Cooldown period](https://docs.aws.amazon.com/autoscaling/application/userguide/step-scaling-policy-overview.html#step-scaling-cooldown) in the *Application Auto Scaling User Guide*.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationautoscaling_scaling_policy#cooldown ApplicationautoscalingScalingPolicy#cooldown}
	Cooldown *float64 `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The aggregation type for the CloudWatch metrics.
	//
	// Valid values are ``Minimum``, ``Maximum``, and ``Average``. If the aggregation type is null, the value is treated as ``Average``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationautoscaling_scaling_policy#metric_aggregation_type ApplicationautoscalingScalingPolicy#metric_aggregation_type}
	MetricAggregationType *string `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// The minimum value to scale by when the adjustment type is ``PercentChangeInCapacity``.
	//
	// For example, suppose that you create a step scaling policy to scale out an Amazon ECS service by 25 percent and you specify a ``MinAdjustmentMagnitude`` of 2. If the service has 4 tasks and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a ``MinAdjustmentMagnitude`` of 2, Application Auto Scaling scales out the service by 2 tasks.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationautoscaling_scaling_policy#min_adjustment_magnitude ApplicationautoscalingScalingPolicy#min_adjustment_magnitude}
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	//
	// At least one step adjustment is required if you are adding a new step scaling policy configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationautoscaling_scaling_policy#step_adjustments ApplicationautoscalingScalingPolicy#step_adjustments}
	StepAdjustments interface{} `field:"optional" json:"stepAdjustments" yaml:"stepAdjustments"`
}

