// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxsnapshot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type FsxSnapshotConfig struct {
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
	// The name of the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_snapshot#name FsxSnapshot#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of the volume that the snapshot is of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_snapshot#volume_id FsxSnapshot#volume_id}
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// A list of Tag values, with a maximum of 50 elements.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.106.0/docs/resources/fsx_snapshot#tags FsxSnapshot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

