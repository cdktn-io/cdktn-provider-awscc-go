// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationsmb


type DatasyncLocationSmbMountOptions struct {
	// The specific SMB version that you want DataSync to use to mount your SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/datasync_location_smb#version DatasyncLocationSmb#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

