// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupbackupplan


type BackupBackupPlanBackupPlanAdvancedBackupSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/backup_backup_plan#backup_options BackupBackupPlan#backup_options}.
	BackupOptions *string `field:"optional" json:"backupOptions" yaml:"backupOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/backup_backup_plan#resource_type BackupBackupPlan#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

