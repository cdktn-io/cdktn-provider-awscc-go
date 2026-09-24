// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxs3accesspointattachment


type FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentity struct {
	// Specifies the FSx for ONTAP user identity type, accepts either UNIX or WINDOWS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_s3_access_point_attachment#type FsxS3AccessPointAttachment#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Specifies the properties of the file system UNIX user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_s3_access_point_attachment#unix_user FsxS3AccessPointAttachment#unix_user}
	UnixUser *FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityUnixUser `field:"optional" json:"unixUser" yaml:"unixUser"`
	// Specifies the properties of the file system Windows user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_s3_access_point_attachment#windows_user FsxS3AccessPointAttachment#windows_user}
	WindowsUser *FsxS3AccessPointAttachmentOntapConfigurationFileSystemIdentityWindowsUser `field:"optional" json:"windowsUser" yaml:"windowsUser"`
}

