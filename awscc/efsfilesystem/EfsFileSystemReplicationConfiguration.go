// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package efsfilesystem


type EfsFileSystemReplicationConfiguration struct {
	// An array of destination objects. Only one destination object is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/efs_file_system#destinations EfsFileSystem#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
}

