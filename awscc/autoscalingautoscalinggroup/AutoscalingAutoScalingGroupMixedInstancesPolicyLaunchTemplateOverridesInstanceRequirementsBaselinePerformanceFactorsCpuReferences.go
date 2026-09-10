// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupMixedInstancesPolicyLaunchTemplateOverridesInstanceRequirementsBaselinePerformanceFactorsCpuReferences struct {
	// The instance family to use as a baseline reference.
	//
	// Make sure that you specify the correct value for the instance family. The instance family is everything before the period (.) in the instance type name. For example, in the instance ``c6i.large``, the instance family is ``c6i``, not ``c6``. For more information, see [Amazon EC2 instance type naming conventions](https://docs.aws.amazon.com/ec2/latest/instancetypes/instance-type-names.html) in *Amazon EC2 Instance Types*.
	//   The following instance types are *not supported* for performance protection.
	//   +   ``c1``
	//   +   ``g3| g3s``
	//   +   ``hpc7g``
	//   +   ``m1| m2``
	//   +   ``mac1 | mac2 | mac2-m1ultra | mac2-m2 | mac2-m2pro``
	//   +   ``p3dn | p4d | p5``
	//   +   ``t1``
	//   +   ``u-12tb1 | u-18tb1 | u-24tb1 | u-3tb1 | u-6tb1 | u-9tb1 | u7i-12tb | u7in-16tb | u7in-24tb | u7in-32tb``
	//
	//  If you performance protection by specifying a supported instance family, the returned instance types will exclude the preceding unsupported instance families.
	//  If you specify an unsupported instance family as a value for baseline performance, the API returns an empty response.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/autoscaling_auto_scaling_group#instance_family AutoscalingAutoScalingGroup#instance_family}
	InstanceFamily *string `field:"optional" json:"instanceFamily" yaml:"instanceFamily"`
}

