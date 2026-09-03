// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOpenZfsConfigurationNfsExportsClientConfigurations struct {
	// A value that specifies who can mount the file system.
	//
	// You can provide a wildcard character (*), an IP address (0.0.0.0), or a CIDR address (192.0.2.0/24). By default, Amazon FSx uses the wildcard character when specifying the client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#clients FsxVolume#clients}
	Clients *string `field:"optional" json:"clients" yaml:"clients"`
	// The configuration object for mounting a Network File System (NFS) file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#options FsxVolume#options}
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
}

