// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package snstopic


type SnsTopicDeliveryStatusLogging struct {
	// The IAM role ARN to be used when logging failed message deliveries in Amazon CloudWatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sns_topic#failure_feedback_role_arn SnsTopic#failure_feedback_role_arn}
	FailureFeedbackRoleArn *string `field:"optional" json:"failureFeedbackRoleArn" yaml:"failureFeedbackRoleArn"`
	// Indicates one of the supported protocols for the Amazon SNS topic.
	//
	// At least one of the other three ``LoggingConfig`` properties is recommend along with ``Protocol``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sns_topic#protocol SnsTopic#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The IAM role ARN to be used when logging successful message deliveries in Amazon CloudWatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sns_topic#success_feedback_role_arn SnsTopic#success_feedback_role_arn}
	SuccessFeedbackRoleArn *string `field:"optional" json:"successFeedbackRoleArn" yaml:"successFeedbackRoleArn"`
	// The percentage of successful message deliveries to be logged in Amazon CloudWatch.
	//
	// Valid percentage values range from 0 to 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sns_topic#success_feedback_sample_rate SnsTopic#success_feedback_sample_rate}
	SuccessFeedbackSampleRate *string `field:"optional" json:"successFeedbackSampleRate" yaml:"successFeedbackSampleRate"`
}

