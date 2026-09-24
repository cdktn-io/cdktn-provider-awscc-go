// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupLifecycleHookSpecificationListStruct struct {
	// The action the Auto Scaling group takes when the lifecycle hook timeout elapses or if an unexpected failure occurs.
	//
	// The default value is ``ABANDON``.
	//  Valid values: ``CONTINUE`` | ``ABANDON``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#default_result AutoscalingAutoScalingGroup#default_result}
	DefaultResult *string `field:"optional" json:"defaultResult" yaml:"defaultResult"`
	// The maximum time, in seconds, that can elapse before the lifecycle hook times out.
	//
	// The range is from ``30`` to ``7200`` seconds. The default value is ``3600`` seconds (1 hour).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#heartbeat_timeout AutoscalingAutoScalingGroup#heartbeat_timeout}
	HeartbeatTimeout *float64 `field:"optional" json:"heartbeatTimeout" yaml:"heartbeatTimeout"`
	// The name of the lifecycle hook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#lifecycle_hook_name AutoscalingAutoScalingGroup#lifecycle_hook_name}
	LifecycleHookName *string `field:"optional" json:"lifecycleHookName" yaml:"lifecycleHookName"`
	// The lifecycle transition.
	//
	// For Auto Scaling groups, there are two major lifecycle transitions.
	//   +  To create a lifecycle hook for scale-out events, specify ``autoscaling:EC2_INSTANCE_LAUNCHING``.
	//   +  To create a lifecycle hook for scale-in events, specify ``autoscaling:EC2_INSTANCE_TERMINATING``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#lifecycle_transition AutoscalingAutoScalingGroup#lifecycle_transition}
	LifecycleTransition *string `field:"optional" json:"lifecycleTransition" yaml:"lifecycleTransition"`
	// Additional information that you want to include any time Amazon EC2 Auto Scaling sends a message to the notification target.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#notification_metadata AutoscalingAutoScalingGroup#notification_metadata}
	NotificationMetadata *string `field:"optional" json:"notificationMetadata" yaml:"notificationMetadata"`
	// The Amazon Resource Name (ARN) of the notification target that Amazon EC2 Auto Scaling sends notifications to when an instance is in a wait state for the lifecycle hook.
	//
	// You can specify an Amazon SNS topic or an Amazon SQS queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#notification_target_arn AutoscalingAutoScalingGroup#notification_target_arn}
	NotificationTargetArn *string `field:"optional" json:"notificationTargetArn" yaml:"notificationTargetArn"`
	// The ARN of the IAM role that allows the Auto Scaling group to publish to the specified notification target.
	//
	// For information about creating this role, see [Prepare to add a lifecycle hook to your Auto Scaling group](https://docs.aws.amazon.com/autoscaling/ec2/userguide/prepare-for-lifecycle-notifications.html) in the *Amazon EC2 Auto Scaling User Guide*.
	//  Valid only if the notification target is an Amazon SNS topic or an Amazon SQS queue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/autoscaling_auto_scaling_group#role_arn AutoscalingAutoScalingGroup#role_arn}
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

