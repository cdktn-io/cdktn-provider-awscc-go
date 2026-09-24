// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisvideostream

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type KinesisvideoStreamConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The number of hours till which Kinesis Video will retain the data in the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisvideo_stream#data_retention_in_hours KinesisvideoStream#data_retention_in_hours}
	DataRetentionInHours *float64 `field:"optional" json:"dataRetentionInHours" yaml:"dataRetentionInHours"`
	// The name of the device that is writing to the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisvideo_stream#device_name KinesisvideoStream#device_name}
	DeviceName *string `field:"optional" json:"deviceName" yaml:"deviceName"`
	// AWS KMS key ID that Kinesis Video Streams uses to encrypt stream data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisvideo_stream#kms_key_id KinesisvideoStream#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The media type of the stream. Consumers of the stream can use this information when processing the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisvideo_stream#media_type KinesisvideoStream#media_type}
	MediaType *string `field:"optional" json:"mediaType" yaml:"mediaType"`
	// The name of the Kinesis Video stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisvideo_stream#name KinesisvideoStream#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Configuration for the storage tier of the Kinesis Video Stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisvideo_stream#stream_storage_configuration KinesisvideoStream#stream_storage_configuration}
	StreamStorageConfiguration *KinesisvideoStreamStreamStorageConfiguration `field:"optional" json:"streamStorageConfiguration" yaml:"streamStorageConfiguration"`
	// An array of key-value pairs associated with the Kinesis Video Stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/kinesisvideo_stream#tags KinesisvideoStream#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

