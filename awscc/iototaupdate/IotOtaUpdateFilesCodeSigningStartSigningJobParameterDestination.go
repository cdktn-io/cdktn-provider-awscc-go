// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateFilesCodeSigningStartSigningJobParameterDestination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#s3_destination IotOtaUpdate#s3_destination}.
	S3Destination *IotOtaUpdateFilesCodeSigningStartSigningJobParameterDestinationS3Destination `field:"optional" json:"s3Destination" yaml:"s3Destination"`
}

