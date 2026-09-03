// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationnfs


type DatasyncLocationNfsMountOptions struct {
	// The specific NFS version that you want DataSync to use to mount your NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/datasync_location_nfs#version DatasyncLocationNfs#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

