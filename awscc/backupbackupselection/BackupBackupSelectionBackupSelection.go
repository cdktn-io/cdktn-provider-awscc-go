// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backupbackupselection


type BackupBackupSelectionBackupSelection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/backup_backup_selection#iam_role_arn BackupBackupSelection#iam_role_arn}.
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/backup_backup_selection#selection_name BackupBackupSelection#selection_name}.
	SelectionName *string `field:"required" json:"selectionName" yaml:"selectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/backup_backup_selection#conditions BackupBackupSelection#conditions}.
	Conditions *BackupBackupSelectionBackupSelectionConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/backup_backup_selection#list_of_tags BackupBackupSelection#list_of_tags}.
	ListOfTags interface{} `field:"optional" json:"listOfTags" yaml:"listOfTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/backup_backup_selection#not_resources BackupBackupSelection#not_resources}.
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/backup_backup_selection#resources BackupBackupSelection#resources}.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
}

