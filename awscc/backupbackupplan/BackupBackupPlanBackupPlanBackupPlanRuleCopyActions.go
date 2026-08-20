// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupbackupplan


type BackupBackupPlanBackupPlanBackupPlanRuleCopyActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/backup_backup_plan#destination_backup_vault_arn BackupBackupPlan#destination_backup_vault_arn}.
	DestinationBackupVaultArn *string `field:"optional" json:"destinationBackupVaultArn" yaml:"destinationBackupVaultArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/backup_backup_plan#lifecycle BackupBackupPlan#lifecycle}.
	Lifecycle *BackupBackupPlanBackupPlanBackupPlanRuleCopyActionsLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
}

