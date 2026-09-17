// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backuplegalhold


type BackupLegalHoldRecoveryPointSelection struct {
	// A date range for filtering recovery points.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/backup_legal_hold#date_range BackupLegalHold#date_range}
	DateRange *BackupLegalHoldRecoveryPointSelectionDateRange `field:"optional" json:"dateRange" yaml:"dateRange"`
	// The resources included in the resource selection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/backup_legal_hold#resource_identifiers BackupLegalHold#resource_identifiers}
	ResourceIdentifiers *[]*string `field:"optional" json:"resourceIdentifiers" yaml:"resourceIdentifiers"`
	// The names of the vaults in which the selected recovery points are contained.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/backup_legal_hold#vault_names BackupLegalHold#vault_names}
	VaultNames *[]*string `field:"optional" json:"vaultNames" yaml:"vaultNames"`
}

