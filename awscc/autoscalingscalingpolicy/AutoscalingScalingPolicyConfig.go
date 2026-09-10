// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingscalingpolicy

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type AutoscalingScalingPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#auto_scaling_group_name AutoscalingScalingPolicy#auto_scaling_group_name}
	AutoScalingGroupName *string `field:"required" json:"autoScalingGroupName" yaml:"autoScalingGroupName"`
	// Specifies how the scaling adjustment is interpreted. The valid values are ChangeInCapacity, ExactCapacity, and PercentChangeInCapacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#adjustment_type AutoscalingScalingPolicy#adjustment_type}
	AdjustmentType *string `field:"optional" json:"adjustmentType" yaml:"adjustmentType"`
	// The duration of the policy's cooldown period, in seconds.
	//
	// When a cooldown period is specified here, it overrides the default cooldown period defined for the Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#cooldown AutoscalingScalingPolicy#cooldown}
	Cooldown *string `field:"optional" json:"cooldown" yaml:"cooldown"`
	// The estimated time, in seconds, until a newly launched instance can contribute to the CloudWatch metrics.
	//
	// If not provided, the default is to use the value from the default cooldown period for the Auto Scaling group. Valid only if the policy type is TargetTrackingScaling or StepScaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#estimated_instance_warmup AutoscalingScalingPolicy#estimated_instance_warmup}
	EstimatedInstanceWarmup *float64 `field:"optional" json:"estimatedInstanceWarmup" yaml:"estimatedInstanceWarmup"`
	// The aggregation type for the CloudWatch metrics.
	//
	// The valid values are Minimum, Maximum, and Average. If the aggregation type is null, the value is treated as Average. Valid only if the policy type is StepScaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#metric_aggregation_type AutoscalingScalingPolicy#metric_aggregation_type}
	MetricAggregationType *string `field:"optional" json:"metricAggregationType" yaml:"metricAggregationType"`
	// The minimum value to scale by when the adjustment type is PercentChangeInCapacity.
	//
	// For example, suppose that you create a step scaling policy to scale out an Auto Scaling group by 25 percent and you specify a MinAdjustmentMagnitude of 2. If the group has 4 instances and the scaling policy is performed, 25 percent of 4 is 1. However, because you specified a MinAdjustmentMagnitude of 2, Amazon EC2 Auto Scaling scales out the group by 2 instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#min_adjustment_magnitude AutoscalingScalingPolicy#min_adjustment_magnitude}
	MinAdjustmentMagnitude *float64 `field:"optional" json:"minAdjustmentMagnitude" yaml:"minAdjustmentMagnitude"`
	// One of the following policy types: TargetTrackingScaling, StepScaling, SimpleScaling (default), PredictiveScaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#policy_type AutoscalingScalingPolicy#policy_type}
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
	// A predictive scaling policy. Includes support for predefined metrics only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#predictive_scaling_configuration AutoscalingScalingPolicy#predictive_scaling_configuration}
	PredictiveScalingConfiguration *AutoscalingScalingPolicyPredictiveScalingConfiguration `field:"optional" json:"predictiveScalingConfiguration" yaml:"predictiveScalingConfiguration"`
	// The amount by which to scale, based on the specified adjustment type.
	//
	// A positive value adds to the current capacity while a negative number removes from the current capacity. For exact capacity, you must specify a positive value. Required if the policy type is SimpleScaling. (Not used with any other policy type.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#scaling_adjustment AutoscalingScalingPolicy#scaling_adjustment}
	ScalingAdjustment *float64 `field:"optional" json:"scalingAdjustment" yaml:"scalingAdjustment"`
	// A set of adjustments that enable you to scale based on the size of the alarm breach.
	//
	// Required if the policy type is StepScaling. (Not used with any other policy type.)
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#step_adjustments AutoscalingScalingPolicy#step_adjustments}
	StepAdjustments interface{} `field:"optional" json:"stepAdjustments" yaml:"stepAdjustments"`
	// A target tracking scaling policy. Includes support for predefined or customized metrics.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_scaling_policy#target_tracking_configuration AutoscalingScalingPolicy#target_tracking_configuration}
	TargetTrackingConfiguration *AutoscalingScalingPolicyTargetTrackingConfiguration `field:"optional" json:"targetTrackingConfiguration" yaml:"targetTrackingConfiguration"`
}

