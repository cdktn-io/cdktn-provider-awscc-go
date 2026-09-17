// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotstream


type IotStreamFiles struct {
	// The file ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_stream#file_id IotStream#file_id}
	FileId *float64 `field:"optional" json:"fileId" yaml:"fileId"`
	// The location of the file in S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_stream#s3_location IotStream#s3_location}
	S3Location *IotStreamFilesS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

