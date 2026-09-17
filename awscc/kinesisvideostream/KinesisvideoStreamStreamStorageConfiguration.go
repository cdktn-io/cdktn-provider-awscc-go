// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisvideostream


type KinesisvideoStreamStreamStorageConfiguration struct {
	// The storage tier for the Kinesis Video Stream. Determines the storage class used for stream data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/kinesisvideo_stream#default_storage_tier KinesisvideoStream#default_storage_tier}
	DefaultStorageTier *string `field:"optional" json:"defaultStorageTier" yaml:"defaultStorageTier"`
}

