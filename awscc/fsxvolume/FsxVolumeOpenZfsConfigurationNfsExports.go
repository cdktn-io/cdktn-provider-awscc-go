// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOpenZfsConfigurationNfsExports struct {
	// The configuration object for mounting a Network File System (NFS) file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#client_configurations FsxVolume#client_configurations}
	ClientConfigurations interface{} `field:"optional" json:"clientConfigurations" yaml:"clientConfigurations"`
}

