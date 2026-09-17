// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverrides struct {
	// The ID of the Amazon Machine Image (AMI) to use for instances launched with this override.
	//
	// When using Instance Refresh with ``ReplaceRootVolume`` strategy, this specifies the AMI for root volume replacement operations.
	//   For ``ReplaceRootVolume`` operations:
	//   +  All overrides in the ``MixedInstancesPolicy`` must specify an ImageId
	//   +  The AMI must contain only a single root volume
	//   +  Root volume replacement doesn't support multi-volume AMIs
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#image_id AutoscalingAutoScalingGroup#image_id}
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// The instance requirements.
	//
	// Amazon EC2 Auto Scaling uses your specified requirements to identify instance types. Then, it uses your On-Demand and Spot allocation strategies to launch instances from these instance types.
	//  You can specify up to four separate sets of instance requirements per Auto Scaling group. This is useful for provisioning instances from different Amazon Machine Images (AMIs) in the same Auto Scaling group. To do this, create the AMIs and create a new launch template for each AMI. Then, create a compatible set of instance requirements for each launch template.
	//   If you specify ``InstanceRequirements``, you can't specify ``InstanceType``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#instance_requirements AutoscalingAutoScalingGroup#instance_requirements}
	InstanceRequirements *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirements `field:"optional" json:"instanceRequirements" yaml:"instanceRequirements"`
	// The instance type, such as ``m3.xlarge``. You must specify an instance type that is supported in your requested Region and Availability Zones. For more information, see [Instance types](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html) in the *Amazon EC2 User Guide*.  You can specify up to 40 instance types per Auto Scaling group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#instance_type AutoscalingAutoScalingGroup#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// Provides a launch template for the specified instance type or set of instance requirements.
	//
	// For example, some instance types might require a launch template with a different AMI. If not provided, Amazon EC2 Auto Scaling uses the launch template that's specified in the ``LaunchTemplate`` definition. For more information, see [Specifying a different launch template for an instance type](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-launch-template-overrides.html) in the *Amazon EC2 Auto Scaling User Guide*.
	//  You can specify up to 20 launch templates per Auto Scaling group. The launch templates specified in the overrides and in the ``LaunchTemplate`` definition count towards this limit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#launch_template_specification AutoscalingAutoScalingGroup#launch_template_specification}
	LaunchTemplateSpecification *AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesLaunchTemplateSpecification `field:"optional" json:"launchTemplateSpecification" yaml:"launchTemplateSpecification"`
	// If you provide a list of instance types to use, you can specify the number of capacity units provided by each instance type in terms of virtual CPUs, memory, storage, throughput, or other relative performance characteristic.
	//
	// When a Spot or On-Demand Instance is launched, the capacity units count toward the desired capacity. Amazon EC2 Auto Scaling launches instances until the desired capacity is totally fulfilled, even if this results in an overage. For example, if there are two units remaining to fulfill capacity, and Amazon EC2 Auto Scaling can only launch an instance with a ``WeightedCapacity`` of five units, the instance is launched, and the desired capacity is exceeded by three units. For more information, see [Configure instance weighting for Amazon EC2 Auto Scaling](https://docs.aws.amazon.com/autoscaling/ec2/userguide/ec2-auto-scaling-mixed-instances-groups-instance-weighting.html) in the *Amazon EC2 Auto Scaling User Guide*. Value must be in the range of 1-999.
	//  If you specify a value for ``WeightedCapacity`` for one instance type, you must specify a value for ``WeightedCapacity`` for all of them.
	//   Every Auto Scaling group has three size parameters (``DesiredCapacity``, ``MaxSize``, and ``MinSize``). Usually, you set these sizes based on a specific number of instances. However, if you configure a mixed instances policy that defines weights for the instance types, you must specify these sizes with the same units that you use for weighting instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/autoscaling_auto_scaling_group#weighted_capacity AutoscalingAutoScalingGroup#weighted_capacity}
	WeightedCapacity *string `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

