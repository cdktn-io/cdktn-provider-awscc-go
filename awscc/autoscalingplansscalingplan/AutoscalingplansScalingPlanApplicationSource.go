// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingplansscalingplan


type AutoscalingplansScalingPlanApplicationSource struct {
	// The Amazon Resource Name (ARN) of a CloudFormation stack. The stack must contain at least one scalable resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#cloudformation_stack_arn AutoscalingplansScalingPlan#cloudformation_stack_arn}
	CloudformationStackArn *string `field:"optional" json:"cloudformationStackArn" yaml:"cloudformationStackArn"`
	// A set of tag filters (keys and values).
	//
	// Each tag filter specified must contain a key with values as optional. Each scaling plan can include up to 50 keys, and each key can include up to 20 values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/autoscalingplans_scaling_plan#tag_filters AutoscalingplansScalingPlan#tag_filters}
	TagFilters interface{} `field:"optional" json:"tagFilters" yaml:"tagFilters"`
}

