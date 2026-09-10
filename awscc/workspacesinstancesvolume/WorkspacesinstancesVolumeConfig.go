// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workspacesinstancesvolume

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type WorkspacesinstancesVolumeConfig struct {
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
	// The Availability Zone in which to create the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesinstances_volume#availability_zone WorkspacesinstancesVolume#availability_zone}
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Indicates whether the volume should be encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesinstances_volume#encrypted WorkspacesinstancesVolume#encrypted}
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
	// The number of I/O operations per second (IOPS).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesinstances_volume#iops WorkspacesinstancesVolume#iops}
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// The identifier of the AWS Key Management Service (AWS KMS) customer master key (CMK) to use for Amazon EBS encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesinstances_volume#kms_key_id WorkspacesinstancesVolume#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The size of the volume, in GiBs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesinstances_volume#size_in_gb WorkspacesinstancesVolume#size_in_gb}
	SizeInGb *float64 `field:"optional" json:"sizeInGb" yaml:"sizeInGb"`
	// The snapshot from which to create the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesinstances_volume#snapshot_id WorkspacesinstancesVolume#snapshot_id}
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// The tags passed to EBS volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesinstances_volume#tag_specifications WorkspacesinstancesVolume#tag_specifications}
	TagSpecifications interface{} `field:"optional" json:"tagSpecifications" yaml:"tagSpecifications"`
	// The throughput to provision for a volume, with a maximum of 1,000 MiB/s.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesinstances_volume#throughput WorkspacesinstancesVolume#throughput}
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// The volume type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/workspacesinstances_volume#volume_type WorkspacesinstancesVolume#volume_type}
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

