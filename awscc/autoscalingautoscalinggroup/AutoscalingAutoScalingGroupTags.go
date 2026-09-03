// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/autoscaling_auto_scaling_group#key AutoscalingAutoScalingGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Set to ``true`` if you want CloudFormation to copy the tag to EC2 instances that are launched as part of the Auto Scaling group.
	//
	// Set to ``false`` if you want the tag attached only to the Auto Scaling group and not copied to any instances launched as part of the Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/autoscaling_auto_scaling_group#propagate_at_launch AutoscalingAutoScalingGroup#propagate_at_launch}
	PropagateAtLaunch interface{} `field:"optional" json:"propagateAtLaunch" yaml:"propagateAtLaunch"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/autoscaling_auto_scaling_group#value AutoscalingAutoScalingGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

