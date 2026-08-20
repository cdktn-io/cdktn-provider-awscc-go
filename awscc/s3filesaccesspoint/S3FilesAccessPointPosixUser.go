// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3filesaccesspoint


type S3FilesAccessPointPosixUser struct {
	// The POSIX group ID used for all file system operations using this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3files_access_point#gid S3FilesAccessPoint#gid}
	Gid *string `field:"optional" json:"gid" yaml:"gid"`
	// Secondary POSIX group IDs used for all file system operations using this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3files_access_point#secondary_gids S3FilesAccessPoint#secondary_gids}
	SecondaryGids *[]*string `field:"optional" json:"secondaryGids" yaml:"secondaryGids"`
	// The POSIX user ID used for all file system operations using this access point.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/s3files_access_point#uid S3FilesAccessPoint#uid}
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

