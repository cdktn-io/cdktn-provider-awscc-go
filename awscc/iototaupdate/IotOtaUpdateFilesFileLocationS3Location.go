// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateFilesFileLocationS3Location struct {
	// The S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#bucket IotOtaUpdate#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#key IotOtaUpdate#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The S3 bucket version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#version IotOtaUpdate#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

