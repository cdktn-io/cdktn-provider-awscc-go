// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package backuplegalhold

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type BackupLegalHoldConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The description of the legal hold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/backup_legal_hold#description BackupLegalHold#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The criteria to assign a set of resources, such as resource types or backup vaults.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/backup_legal_hold#recovery_point_selection BackupLegalHold#recovery_point_selection}
	RecoveryPointSelection *BackupLegalHoldRecoveryPointSelection `field:"required" json:"recoveryPointSelection" yaml:"recoveryPointSelection"`
	// The title of the legal hold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/backup_legal_hold#title BackupLegalHold#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Optional tags to include.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/backup_legal_hold#tags BackupLegalHold#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

