// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package smsvoiceconfigurationset


type SmsvoiceConfigurationSetEventDestinations struct {
	// An object that contains IamRoleArn and LogGroupArn associated with an Amazon CloudWatch event destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/smsvoice_configuration_set#cloudwatch_logs_destination SmsvoiceConfigurationSet#cloudwatch_logs_destination}
	CloudwatchLogsDestination *SmsvoiceConfigurationSetEventDestinationsCloudwatchLogsDestination `field:"optional" json:"cloudwatchLogsDestination" yaml:"cloudwatchLogsDestination"`
	// When set to true events will be logged. By default this is set to true.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/smsvoice_configuration_set#enabled SmsvoiceConfigurationSet#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The name that identifies the event destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/smsvoice_configuration_set#event_destination_name SmsvoiceConfigurationSet#event_destination_name}
	EventDestinationName *string `field:"optional" json:"eventDestinationName" yaml:"eventDestinationName"`
	// An object that contains IamRoleArn and DeliveryStreamArn associated with an Amazon Kinesis Firehose event destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/smsvoice_configuration_set#kinesis_firehose_destination SmsvoiceConfigurationSet#kinesis_firehose_destination}
	KinesisFirehoseDestination *SmsvoiceConfigurationSetEventDestinationsKinesisFirehoseDestination `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// An array of event types that determine which events to log.
	//
	// If 'ALL' is used, then AWS End User Messaging SMS and Voice logs every event type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/smsvoice_configuration_set#matching_event_types SmsvoiceConfigurationSet#matching_event_types}
	MatchingEventTypes *[]*string `field:"optional" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// An object that contains SNS TopicArn event destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/smsvoice_configuration_set#sns_destination SmsvoiceConfigurationSet#sns_destination}
	SnsDestination *SmsvoiceConfigurationSetEventDestinationsSnsDestination `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

