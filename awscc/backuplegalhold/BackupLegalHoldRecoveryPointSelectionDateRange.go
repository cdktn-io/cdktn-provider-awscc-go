// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backuplegalhold


type BackupLegalHoldRecoveryPointSelectionDateRange struct {
	// The beginning date, inclusive. ISO 8601 date-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/backup_legal_hold#from_date BackupLegalHold#from_date}
	FromDate *string `field:"optional" json:"fromDate" yaml:"fromDate"`
	// The end date, inclusive. ISO 8601 date-time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/backup_legal_hold#to_date BackupLegalHold#to_date}
	ToDate *string `field:"optional" json:"toDate" yaml:"toDate"`
}

