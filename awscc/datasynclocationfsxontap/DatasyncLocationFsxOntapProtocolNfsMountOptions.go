// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxontap


type DatasyncLocationFsxOntapProtocolNfsMountOptions struct {
	// The specific NFS version that you want DataSync to use to mount your NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datasync_location_fsx_ontap#version DatasyncLocationFsxOntap#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

