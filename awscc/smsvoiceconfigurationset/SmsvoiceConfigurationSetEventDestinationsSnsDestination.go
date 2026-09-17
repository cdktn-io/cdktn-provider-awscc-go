// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoiceconfigurationset


type SmsvoiceConfigurationSetEventDestinationsSnsDestination struct {
	// The Amazon Resource Name (ARN) of the Amazon SNS topic that you want to publish events to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/smsvoice_configuration_set#topic_arn SmsvoiceConfigurationSet#topic_arn}
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

