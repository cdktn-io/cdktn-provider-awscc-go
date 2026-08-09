// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mskchannel


type MskChannelLoggingInfoS3 struct {
	// The name of the S3 bucket for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#bucket MskChannel#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// Whether S3 logging is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#enabled MskChannel#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The S3 prefix for log delivery.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/msk_channel#prefix MskChannel#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

