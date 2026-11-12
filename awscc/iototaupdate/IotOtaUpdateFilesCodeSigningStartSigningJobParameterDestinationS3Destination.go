// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateFilesCodeSigningStartSigningJobParameterDestinationS3Destination struct {
	// The S3 bucket that contains the updated firmware.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#bucket IotOtaUpdate#bucket}
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The S3 prefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#prefix IotOtaUpdate#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

