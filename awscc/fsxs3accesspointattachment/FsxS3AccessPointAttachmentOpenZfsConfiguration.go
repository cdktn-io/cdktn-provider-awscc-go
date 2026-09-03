// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentOpenZfsConfiguration struct {
	// The file system identity used to authorize file access requests made using the S3 access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_s3_access_point_attachment#file_system_identity FsxS3AccessPointAttachment#file_system_identity}
	FileSystemIdentity *FsxS3AccessPointAttachmentOpenZfsConfigurationFileSystemIdentity `field:"optional" json:"fileSystemIdentity" yaml:"fileSystemIdentity"`
	// The ID of the FSx for OpenZFS volume that the S3 access point is attached to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_s3_access_point_attachment#volume_id FsxS3AccessPointAttachment#volume_id}
	VolumeId *string `field:"optional" json:"volumeId" yaml:"volumeId"`
}

