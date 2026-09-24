// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesischannel


type KinesisChannelLoggingConfiguration struct {
	// CloudWatch Logs configuration block. When provided, controls whether and where the channel writes operational logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesis_channel#cloudwatch_logs KinesisChannel#cloudwatch_logs}
	CloudwatchLogs *KinesisChannelLoggingConfigurationCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
}

