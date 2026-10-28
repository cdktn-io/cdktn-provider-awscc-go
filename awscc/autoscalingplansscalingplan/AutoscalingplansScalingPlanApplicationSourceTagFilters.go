// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanApplicationSourceTagFilters struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#key AutoscalingplansScalingPlan#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag values (0 to 20).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#values AutoscalingplansScalingPlan#values}
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

