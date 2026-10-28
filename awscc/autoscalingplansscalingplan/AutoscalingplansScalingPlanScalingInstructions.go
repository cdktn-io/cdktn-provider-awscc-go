// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanScalingInstructions struct {
	// The maximum capacity of the resource.
	//
	// The exception to this upper limit is if you specify a non-default setting for PredictiveScalingMaxCapacityBehavior.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#max_capacity AutoscalingplansScalingPlan#max_capacity}
	MaxCapacity *float64 `field:"required" json:"maxCapacity" yaml:"maxCapacity"`
	// The minimum capacity of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#min_capacity AutoscalingplansScalingPlan#min_capacity}
	MinCapacity *float64 `field:"required" json:"minCapacity" yaml:"minCapacity"`
	// The ID of the resource. This string consists of the resource type and unique identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#resource_id AutoscalingplansScalingPlan#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// The scalable dimension associated with the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#scalable_dimension AutoscalingplansScalingPlan#scalable_dimension}
	ScalableDimension *string `field:"required" json:"scalableDimension" yaml:"scalableDimension"`
	// The namespace of the AWS service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#service_namespace AutoscalingplansScalingPlan#service_namespace}
	ServiceNamespace *string `field:"required" json:"serviceNamespace" yaml:"serviceNamespace"`
	// The target tracking configurations (up to 10).
	//
	// Each of these structures must specify a unique scaling metric and a target value for the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#target_tracking_configurations AutoscalingplansScalingPlan#target_tracking_configurations}
	TargetTrackingConfigurations interface{} `field:"required" json:"targetTrackingConfigurations" yaml:"targetTrackingConfigurations"`
	// The customized load metric to use for predictive scaling.
	//
	// This property or a PredefinedLoadMetricSpecification is required when configuring predictive scaling, and cannot be used otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#customized_load_metric_specification AutoscalingplansScalingPlan#customized_load_metric_specification}
	CustomizedLoadMetricSpecification *AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecification `field:"optional" json:"customizedLoadMetricSpecification" yaml:"customizedLoadMetricSpecification"`
	// Controls whether dynamic scaling is disabled.
	//
	// When dynamic scaling is enabled, AWS Auto Scaling creates target tracking scaling policies based on the specified target tracking configurations. The default is enabled (false).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#disable_dynamic_scaling AutoscalingplansScalingPlan#disable_dynamic_scaling}
	DisableDynamicScaling interface{} `field:"optional" json:"disableDynamicScaling" yaml:"disableDynamicScaling"`
	// The predefined load metric to use for predictive scaling.
	//
	// This property or a CustomizedLoadMetricSpecification is required when configuring predictive scaling, and cannot be used otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#predefined_load_metric_specification AutoscalingplansScalingPlan#predefined_load_metric_specification}
	PredefinedLoadMetricSpecification *AutoscalingplansScalingPlanScalingInstructionsPredefinedLoadMetricSpecification `field:"optional" json:"predefinedLoadMetricSpecification" yaml:"predefinedLoadMetricSpecification"`
	// Defines the behavior that should be applied if the forecast capacity approaches or exceeds the maximum capacity specified for the resource.
	//
	// The default value is SetForecastCapacityToMaxCapacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#predictive_scaling_max_capacity_behavior AutoscalingplansScalingPlan#predictive_scaling_max_capacity_behavior}
	PredictiveScalingMaxCapacityBehavior *string `field:"optional" json:"predictiveScalingMaxCapacityBehavior" yaml:"predictiveScalingMaxCapacityBehavior"`
	// The size of the capacity buffer to use when the forecast capacity is close to or exceeds the maximum capacity.
	//
	// The value is specified as a percentage relative to the forecast capacity. For example, if the buffer is 10, this means a 10 percent buffer. The range is 1-100. Required if PredictiveScalingMaxCapacityBehavior is set to SetMaxCapacityAboveForecastCapacity, and cannot be used otherwise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#predictive_scaling_max_capacity_buffer AutoscalingplansScalingPlan#predictive_scaling_max_capacity_buffer}
	PredictiveScalingMaxCapacityBuffer *float64 `field:"optional" json:"predictiveScalingMaxCapacityBuffer" yaml:"predictiveScalingMaxCapacityBuffer"`
	// The predictive scaling mode.
	//
	// The default value is ForecastAndScale. Otherwise, AWS Auto Scaling forecasts capacity but does not apply any scheduled scaling actions based on the capacity forecast.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#predictive_scaling_mode AutoscalingplansScalingPlan#predictive_scaling_mode}
	PredictiveScalingMode *string `field:"optional" json:"predictiveScalingMode" yaml:"predictiveScalingMode"`
	// Controls whether a resource's externally created scaling policies are deleted and new target tracking scaling policies created.
	//
	// The default value is KeepExternalPolicies. Valid only when configuring dynamic scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#scaling_policy_update_behavior AutoscalingplansScalingPlan#scaling_policy_update_behavior}
	ScalingPolicyUpdateBehavior *string `field:"optional" json:"scalingPolicyUpdateBehavior" yaml:"scalingPolicyUpdateBehavior"`
	// The amount of time, in seconds, to buffer the run time of scheduled scaling actions when scaling out.
	//
	// The value must be less than the forecast interval duration of 3600 seconds (60 minutes). The default is 300 seconds. Valid only when configuring predictive scaling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#scheduled_action_buffer_time AutoscalingplansScalingPlan#scheduled_action_buffer_time}
	ScheduledActionBufferTime *float64 `field:"optional" json:"scheduledActionBufferTime" yaml:"scheduledActionBufferTime"`
}

