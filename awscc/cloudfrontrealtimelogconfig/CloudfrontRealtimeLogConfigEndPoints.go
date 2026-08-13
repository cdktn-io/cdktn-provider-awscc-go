// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudfrontrealtimelogconfig


type CloudfrontRealtimeLogConfigEndPoints struct {
	// Contains information about the Amazon Kinesis data stream where you are sending real-time log data in a real-time log configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudfront_realtime_log_config#kinesis_stream_config CloudfrontRealtimeLogConfig#kinesis_stream_config}
	KinesisStreamConfig *CloudfrontRealtimeLogConfigEndPointsKinesisStreamConfig `field:"required" json:"kinesisStreamConfig" yaml:"kinesisStreamConfig"`
	// The type of data stream where you are sending real-time log data. The only valid value is ``Kinesis``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/cloudfront_realtime_log_config#stream_type CloudfrontRealtimeLogConfig#stream_type}
	StreamType *string `field:"required" json:"streamType" yaml:"streamType"`
}

