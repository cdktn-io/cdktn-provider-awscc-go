// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fsxvolume


type FsxVolumeOpenZfsConfigurationOriginSnapshot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_volume#copy_strategy FsxVolume#copy_strategy}.
	CopyStrategy *string `field:"optional" json:"copyStrategy" yaml:"copyStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fsx_volume#snapshot_arn FsxVolume#snapshot_arn}.
	SnapshotArn *string `field:"optional" json:"snapshotArn" yaml:"snapshotArn"`
}

