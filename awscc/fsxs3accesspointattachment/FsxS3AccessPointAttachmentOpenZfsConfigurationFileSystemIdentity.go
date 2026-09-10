// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentOpenZfsConfigurationFileSystemIdentity struct {
	// Specifies the UID and GIDs of the file system POSIX user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_s3_access_point_attachment#posix_user FsxS3AccessPointAttachment#posix_user}
	PosixUser *FsxS3AccessPointAttachmentOpenZfsConfigurationFileSystemIdentityPosixUser `field:"optional" json:"posixUser" yaml:"posixUser"`
	// Specifies the FSx for OpenZFS user identity type, accepts only POSIX.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/fsx_s3_access_point_attachment#type FsxS3AccessPointAttachment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

