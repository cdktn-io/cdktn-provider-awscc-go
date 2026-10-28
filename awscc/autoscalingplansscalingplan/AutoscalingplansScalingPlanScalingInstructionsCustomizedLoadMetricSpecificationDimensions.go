// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanScalingInstructionsCustomizedLoadMetricSpecificationDimensions struct {
	// The name of the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#name AutoscalingplansScalingPlan#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The value of the dimension.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#value AutoscalingplansScalingPlan#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

