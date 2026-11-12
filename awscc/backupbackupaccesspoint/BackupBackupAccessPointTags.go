// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupbackupaccesspoint


type BackupBackupAccessPointTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/backup_backup_access_point#key BackupBackupAccessPoint#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/backup_backup_access_point#value BackupBackupAccessPoint#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

