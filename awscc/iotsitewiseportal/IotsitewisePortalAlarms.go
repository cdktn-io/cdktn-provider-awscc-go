// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseportal


type IotsitewisePortalAlarms struct {
	// The ARN of the IAM role that allows the alarm to perform actions and access AWS resources and services, such as AWS IoT Events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_portal#alarm_role_arn IotsitewisePortal#alarm_role_arn}
	AlarmRoleArn *string `field:"optional" json:"alarmRoleArn" yaml:"alarmRoleArn"`
	// The ARN of the AWS Lambda function that manages alarm notifications.
	//
	// For more information, see Managing alarm notifications in the AWS IoT Events Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/iotsitewise_portal#notification_lambda_arn IotsitewisePortal#notification_lambda_arn}
	NotificationLambdaArn *string `field:"optional" json:"notificationLambdaArn" yaml:"notificationLambdaArn"`
}

