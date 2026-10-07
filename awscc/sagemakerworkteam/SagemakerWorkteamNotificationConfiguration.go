// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam


type SagemakerWorkteamNotificationConfiguration struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications should be published.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/sagemaker_workteam#notification_topic_arn SagemakerWorkteam#notification_topic_arn}
	NotificationTopicArn *string `field:"optional" json:"notificationTopicArn" yaml:"notificationTopicArn"`
}

