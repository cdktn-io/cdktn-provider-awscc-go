// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationautoscalingscalingpolicy


type ApplicationautoscalingScalingPolicyPredictiveScalingPolicyConfiguration struct {
	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity.
	//
	// Defaults to ``HonorMaxCapacity`` if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#max_capacity_breach_behavior ApplicationautoscalingScalingPolicy#max_capacity_breach_behavior}
	MaxCapacityBreachBehavior *string `field:"optional" json:"maxCapacityBreachBehavior" yaml:"maxCapacityBreachBehavior"`
	// The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	//
	// The value is specified as a percentage relative to the forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer, such that if the forecast capacity is 50, and the maximum capacity is 40, then the effective maximum capacity is 55.
	//  Required if the ``MaxCapacityBreachBehavior`` property is set to ``IncreaseMaxCapacity``, and cannot be used otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#max_capacity_buffer ApplicationautoscalingScalingPolicy#max_capacity_buffer}
	MaxCapacityBuffer *float64 `field:"optional" json:"maxCapacityBuffer" yaml:"maxCapacityBuffer"`
	// This structure includes the metrics and target utilization to use for predictive scaling.
	//
	// This is an array, but we currently only support a single metric specification. That is, you can specify a target value and a single metric pair, or a target value and one scaling metric and one load metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#metric_specifications ApplicationautoscalingScalingPolicy#metric_specifications}
	MetricSpecifications interface{} `field:"optional" json:"metricSpecifications" yaml:"metricSpecifications"`
	// The predictive scaling mode. Defaults to ``ForecastOnly`` if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#mode ApplicationautoscalingScalingPolicy#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// The amount of time, in seconds, that the start time can be advanced.
	//
	// The value must be less than the forecast interval duration of 3600 seconds (60 minutes). Defaults to 300 seconds if not specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationautoscaling_scaling_policy#scheduling_buffer_time ApplicationautoscalingScalingPolicy#scheduling_buffer_time}
	SchedulingBufferTime *float64 `field:"optional" json:"schedulingBufferTime" yaml:"schedulingBufferTime"`
}

