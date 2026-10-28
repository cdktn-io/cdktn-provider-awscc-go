// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecification struct {
	// The dimensions of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#dimensions AutoscalingplansScalingPlan#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The name of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#metric_name AutoscalingplansScalingPlan#metric_name}
	MetricName *string `field:"optional" json:"metricName" yaml:"metricName"`
	// The namespace of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#namespace AutoscalingplansScalingPlan#namespace}
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// The statistic of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#statistic AutoscalingplansScalingPlan#statistic}
	Statistic *string `field:"optional" json:"statistic" yaml:"statistic"`
	// The unit of the metric.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#unit AutoscalingplansScalingPlan#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

