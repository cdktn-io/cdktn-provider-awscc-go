// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxopenzfs


type DatasyncLocationFsxOpenZfsProtocolNfs struct {
	// The NFS mount options that DataSync can use to mount your NFS share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/datasync_location_fsx_open_zfs#mount_options DatasyncLocationFsxOpenZfs#mount_options}
	MountOptions *DatasyncLocationFsxOpenZfsProtocolNfsMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

