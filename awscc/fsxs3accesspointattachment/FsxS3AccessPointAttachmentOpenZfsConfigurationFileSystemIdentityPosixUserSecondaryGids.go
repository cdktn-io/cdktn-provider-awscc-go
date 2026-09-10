// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentOpenZfsConfigurationFileSystemIdentityPosixUserSecondaryGids struct {
	// The GID of the file system user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_s3_access_point_attachment#gid FsxS3AccessPointAttachment#gid}
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
}

