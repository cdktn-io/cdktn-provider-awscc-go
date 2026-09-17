// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsdbsnapshot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsDbSnapshotConfig struct {
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
	// The identifier of the DB instance that you want to create the snapshot of.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_db_snapshot#db_instance_identifier RdsDbSnapshot#db_instance_identifier}
	DbInstanceIdentifier *string `field:"required" json:"dbInstanceIdentifier" yaml:"dbInstanceIdentifier"`
	// The identifier for the DB snapshot.
	//
	// Must contain from 1 to 255 letters, numbers, or hyphens. First character must be a letter. Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_db_snapshot#db_snapshot_identifier RdsDbSnapshot#db_snapshot_identifier}
	DbSnapshotIdentifier *string `field:"required" json:"dbSnapshotIdentifier" yaml:"dbSnapshotIdentifier"`
	// The tags to be assigned to the DB snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/rds_db_snapshot#tags RdsDbSnapshot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

