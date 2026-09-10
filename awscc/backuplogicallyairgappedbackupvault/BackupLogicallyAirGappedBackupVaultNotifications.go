// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backuplogicallyairgappedbackupvault


type BackupLogicallyAirGappedBackupVaultNotifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/backup_logically_air_gapped_backup_vault#backup_vault_events BackupLogicallyAirGappedBackupVault#backup_vault_events}.
	BackupVaultEvents *[]*string `field:"optional" json:"backupVaultEvents" yaml:"backupVaultEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/backup_logically_air_gapped_backup_vault#sns_topic_arn BackupLogicallyAirGappedBackupVault#sns_topic_arn}.
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

