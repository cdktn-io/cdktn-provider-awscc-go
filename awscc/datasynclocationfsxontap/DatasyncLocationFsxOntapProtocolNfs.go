// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxontap


type DatasyncLocationFsxOntapProtocolNfs struct {
	// The NFS mount options that DataSync can use to mount your NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datasync_location_fsx_ontap#mount_options DatasyncLocationFsxOntap#mount_options}
	MountOptions *DatasyncLocationFsxOntapProtocolNfsMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

