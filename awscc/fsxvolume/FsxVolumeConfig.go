// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FsxVolumeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The name of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#name FsxVolume#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Specifies the ID of the volume backup to use to create a new volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#backup_id FsxVolume#backup_id}
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// The configuration of an Amazon FSx for NetApp ONTAP volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#ontap_configuration FsxVolume#ontap_configuration}
	OntapConfiguration *FsxVolumeOntapConfiguration `field:"optional" json:"ontapConfiguration" yaml:"ontapConfiguration"`
	// The configuration of an Amazon FSx for OpenZFS volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#open_zfs_configuration FsxVolume#open_zfs_configuration}
	OpenZfsConfiguration *FsxVolumeOpenZfsConfiguration `field:"optional" json:"openZfsConfiguration" yaml:"openZfsConfiguration"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#tags FsxVolume#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/fsx_volume#volume_type FsxVolume#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

