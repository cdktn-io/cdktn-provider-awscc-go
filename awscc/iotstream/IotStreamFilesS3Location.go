// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotstream


type IotStreamFilesS3Location struct {
	// The S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_stream#bucket IotStream#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_stream#key IotStream#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The S3 bucket version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_stream#version IotStream#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

