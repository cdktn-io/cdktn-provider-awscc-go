// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanScalingInstructionsPredefinedLoadMetricSpecification struct {
	// The metric type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#predefined_load_metric_type AutoscalingplansScalingPlan#predefined_load_metric_type}
	PredefinedLoadMetricType *string `field:"optional" json:"predefinedLoadMetricType" yaml:"predefinedLoadMetricType"`
	// Identifies the resource associated with the metric type.
	//
	// You can't specify a resource label unless the metric type is ALBTargetGroupRequestCount and there is a target group for an Application Load Balancer attached to the Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#resource_label AutoscalingplansScalingPlan#resource_label}
	ResourceLabel *string `field:"optional" json:"resourceLabel" yaml:"resourceLabel"`
}

