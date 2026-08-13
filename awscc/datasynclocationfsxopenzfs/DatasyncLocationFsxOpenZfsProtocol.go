// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datasynclocationfsxopenzfs


type DatasyncLocationFsxOpenZfsProtocol struct {
	// FSx OpenZFS file system NFS protocol information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/datasync_location_fsx_open_zfs#nfs DatasyncLocationFsxOpenZfs#nfs}
	Nfs *DatasyncLocationFsxOpenZfsProtocolNfs `field:"optional" json:"nfs" yaml:"nfs"`
}

