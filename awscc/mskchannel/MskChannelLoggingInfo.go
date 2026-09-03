// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelLoggingInfo struct {
	// CloudWatch Logs log destination details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_channel#cloudwatch_logs MskChannel#cloudwatch_logs}
	CloudwatchLogs *MskChannelLoggingInfoCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Firehose log destination details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_channel#firehose MskChannel#firehose}
	Firehose *MskChannelLoggingInfoFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// S3 log destination details.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/msk_channel#s3 MskChannel#s3}
	S3 *MskChannelLoggingInfoS3 `field:"optional" json:"s3" yaml:"s3"`
}

