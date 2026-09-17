// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package redshiftsnapshot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RedshiftSnapshotConfig struct {
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
	// The cluster identifier for which you want a snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_snapshot#cluster_identifier RedshiftSnapshot#cluster_identifier}
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// A unique identifier for the snapshot that you are requesting.
	//
	// This identifier must be unique for all snapshots within the Amazon Web Services account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_snapshot#snapshot_identifier RedshiftSnapshot#snapshot_identifier}
	SnapshotIdentifier *string `field:"required" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// The number of days that a manual snapshot is retained.
	//
	// If the value is -1, the manual snapshot is retained indefinitely. The value must be either -1 or an integer between 1 and 3653.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_snapshot#manual_snapshot_retention_period RedshiftSnapshot#manual_snapshot_retention_period}
	ManualSnapshotRetentionPeriod *float64 `field:"optional" json:"manualSnapshotRetentionPeriod" yaml:"manualSnapshotRetentionPeriod"`
	// A list of tag instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/redshift_snapshot#tags RedshiftSnapshot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

