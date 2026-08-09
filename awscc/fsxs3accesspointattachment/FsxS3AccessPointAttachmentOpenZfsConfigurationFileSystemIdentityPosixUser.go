// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentOpenZfsConfigurationFileSystemIdentityPosixUser struct {
	// The GID of the file system user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/fsx_s3_access_point_attachment#gid FsxS3AccessPointAttachment#gid}
	Gid *float64 `field:"optional" json:"gid" yaml:"gid"`
	// The list of secondary GIDs for the file system user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/fsx_s3_access_point_attachment#secondary_gids FsxS3AccessPointAttachment#secondary_gids}
	SecondaryGids interface{} `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
	// The UID of the file system user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/fsx_s3_access_point_attachment#uid FsxS3AccessPointAttachment#uid}
	Uid *float64 `field:"optional" json:"uid" yaml:"uid"`
}

