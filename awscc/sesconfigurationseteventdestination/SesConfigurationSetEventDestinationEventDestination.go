// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sesconfigurationseteventdestination


type SesConfigurationSetEventDestinationEventDestination struct {
	// The type of email sending events, send, reject, bounce, complaint, delivery, open, click, renderingFailure, deliveryDelay, and subscription.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_configuration_set_event_destination#matching_event_types SesConfigurationSetEventDestination#matching_event_types}
	MatchingEventTypes *[]*string `field:"required" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// An object that contains the names, default values, and sources of the dimensions associated with an Amazon CloudWatch event destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_configuration_set_event_destination#cloudwatch_destination SesConfigurationSetEventDestination#cloudwatch_destination}
	CloudwatchDestination *SesConfigurationSetEventDestinationEventDestinationCloudwatchDestination `field:"optional" json:"cloudwatchDestination" yaml:"cloudwatchDestination"`
	// Sets whether Amazon SES publishes events to this destination when you send an email with the associated configuration set.
	//
	// Set to true to enable publishing to this destination; set to false to prevent publishing to this destination. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_configuration_set_event_destination#enabled SesConfigurationSetEventDestination#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// An object that contains Event bus ARN associated with the event bridge destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_configuration_set_event_destination#event_bridge_destination SesConfigurationSetEventDestination#event_bridge_destination}
	EventBridgeDestination *SesConfigurationSetEventDestinationEventDestinationEventBridgeDestination `field:"optional" json:"eventBridgeDestination" yaml:"eventBridgeDestination"`
	// An object that contains the delivery stream ARN and the IAM role ARN associated with an Amazon Kinesis Firehose event destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_configuration_set_event_destination#kinesis_firehose_destination SesConfigurationSetEventDestination#kinesis_firehose_destination}
	KinesisFirehoseDestination *SesConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// The name of the event destination set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_configuration_set_event_destination#name SesConfigurationSetEventDestination#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An object that contains SNS topic ARN associated event destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ses_configuration_set_event_destination#sns_destination SesConfigurationSetEventDestination#sns_destination}
	SnsDestination *SesConfigurationSetEventDestinationEventDestinationSnsDestination `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

