// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupbackupvault


type BackupBackupVaultLockConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/backup_backup_vault#changeable_for_days BackupBackupVault#changeable_for_days}.
	ChangeableForDays *float64 `field:"optional" json:"changeableForDays" yaml:"changeableForDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/backup_backup_vault#max_retention_days BackupBackupVault#max_retention_days}.
	MaxRetentionDays *float64 `field:"optional" json:"maxRetentionDays" yaml:"maxRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/backup_backup_vault#min_retention_days BackupBackupVault#min_retention_days}.
	MinRetentionDays *float64 `field:"optional" json:"minRetentionDays" yaml:"minRetentionDays"`
}

