// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisstream


type KinesisStreamStreamModeDetails struct {
	// The mode of the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/kinesis_stream#stream_mode KinesisStream#stream_mode}
	StreamMode *string `field:"optional" json:"streamMode" yaml:"streamMode"`
}

