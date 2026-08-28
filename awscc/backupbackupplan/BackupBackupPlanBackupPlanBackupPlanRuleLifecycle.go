// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupbackupplan


type BackupBackupPlanBackupPlanBackupPlanRuleLifecycle struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/backup_backup_plan#delete_after_days BackupBackupPlan#delete_after_days}.
	DeleteAfterDays *float64 `field:"optional" json:"deleteAfterDays" yaml:"deleteAfterDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/backup_backup_plan#move_to_cold_storage_after_days BackupBackupPlan#move_to_cold_storage_after_days}.
	MoveToColdStorageAfterDays *float64 `field:"optional" json:"moveToColdStorageAfterDays" yaml:"moveToColdStorageAfterDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/backup_backup_plan#opt_in_to_archive_for_supported_resources BackupBackupPlan#opt_in_to_archive_for_supported_resources}.
	OptInToArchiveForSupportedResources interface{} `field:"optional" json:"optInToArchiveForSupportedResources" yaml:"optInToArchiveForSupportedResources"`
}

