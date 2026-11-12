// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iototaupdate


type IotOtaUpdateFilesFileLocationStream struct {
	// The ID of a file associated with a stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#file_id IotOtaUpdate#file_id}
	FileId *float64 `field:"optional" json:"fileId" yaml:"fileId"`
	// The stream ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/iot_ota_update#stream_id IotOtaUpdate#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

