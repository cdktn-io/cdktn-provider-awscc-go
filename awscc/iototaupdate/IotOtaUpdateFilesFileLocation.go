// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateFilesFileLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#s3_location IotOtaUpdate#s3_location}.
	S3Location *IotOtaUpdateFilesFileLocationS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#stream IotOtaUpdate#stream}.
	Stream *IotOtaUpdateFilesFileLocationStream `field:"optional" json:"stream" yaml:"stream"`
}

