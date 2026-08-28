// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lightsaildisksnapshot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type LightsailDiskSnapshotConfig struct {
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
	// The name of the source disk from which the snapshot was created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_disk_snapshot#disk_name LightsailDiskSnapshot#disk_name}
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// The name of the disk snapshot (e.g., my-disk-snapshot).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_disk_snapshot#disk_snapshot_name LightsailDiskSnapshot#disk_snapshot_name}
	DiskSnapshotName *string `field:"required" json:"diskSnapshotName" yaml:"diskSnapshotName"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/lightsail_disk_snapshot#tags LightsailDiskSnapshot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

