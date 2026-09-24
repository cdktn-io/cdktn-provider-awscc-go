// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingscalingpolicy


type AutoscalingScalingPolicyPredictiveScalingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_scaling_policy#max_capacity_breach_behavior AutoscalingScalingPolicy#max_capacity_breach_behavior}.
	MaxCapacityBreachBehavior *string `field:"optional" json:"maxCapacityBreachBehavior" yaml:"maxCapacityBreachBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_scaling_policy#max_capacity_buffer AutoscalingScalingPolicy#max_capacity_buffer}.
	MaxCapacityBuffer *float64 `field:"optional" json:"maxCapacityBuffer" yaml:"maxCapacityBuffer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_scaling_policy#metric_specifications AutoscalingScalingPolicy#metric_specifications}.
	MetricSpecifications interface{} `field:"optional" json:"metricSpecifications" yaml:"metricSpecifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_scaling_policy#mode AutoscalingScalingPolicy#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_scaling_policy#scheduling_buffer_time AutoscalingScalingPolicy#scheduling_buffer_time}.
	SchedulingBufferTime *float64 `field:"optional" json:"schedulingBufferTime" yaml:"schedulingBufferTime"`
}

