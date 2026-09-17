// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesLaunchTemplateSpecification struct {
	// The ID of the launch template.  You must specify the ``LaunchTemplateID`` or the ``LaunchTemplateName``, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#launch_template_id AutoscalingAutoScalingGroup#launch_template_id}
	LaunchTemplateId *string `field:"optional" json:"launchTemplateId" yaml:"launchTemplateId"`
	// The name of the launch template.  You must specify the ``LaunchTemplateName`` or the ``LaunchTemplateID``, but not both.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#launch_template_name AutoscalingAutoScalingGroup#launch_template_name}
	LaunchTemplateName *string `field:"optional" json:"launchTemplateName" yaml:"launchTemplateName"`
	// The version number of the launch template.
	//
	// Specifying ``$Latest`` or ``$Default`` for the template version number is not supported. However, you can specify ``LatestVersionNumber`` or ``DefaultVersionNumber`` using the ``Fn::GetAtt`` intrinsic function. For more information, see [Fn::GetAtt](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-getatt.html).
	//   For an example of using the ``Fn::GetAtt`` function, see the [Examples](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-autoscaling-autoscalinggroup.html#aws-resource-autoscaling-autoscalinggroup--examples) section of the ``AWS::AutoScaling::AutoScalingGroup`` resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#version AutoscalingAutoScalingGroup#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

