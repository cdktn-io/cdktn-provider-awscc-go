// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package autoscalingautoscalinggroup


type AutoscalingAutoScalingGroupNotificationConfigurations struct {
	// A list of event types that send a notification.
	//
	// Event types can include any of the following types.
	//  *Allowed values*:
	//   +   ``autoscaling:EC2_INSTANCE_LAUNCH``
	//   +   ``autoscaling:EC2_INSTANCE_LAUNCH_ERROR``
	//   +   ``autoscaling:EC2_INSTANCE_TERMINATE``
	//   +   ``autoscaling:EC2_INSTANCE_TERMINATE_ERROR``
	//   +   ``autoscaling:TEST_NOTIFICATION``
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/autoscaling_auto_scaling_group#notification_types AutoscalingAutoScalingGroup#notification_types}
	NotificationTypes *[]*string `field:"optional" json:"notificationTypes" yaml:"notificationTypes"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/autoscaling_auto_scaling_group#topic_arn AutoscalingAutoScalingGroup#topic_arn}
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

