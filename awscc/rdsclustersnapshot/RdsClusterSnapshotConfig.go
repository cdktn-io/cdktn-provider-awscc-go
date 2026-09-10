// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rdsclustersnapshot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type RdsClusterSnapshotConfig struct {
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
	// The identifier of the DB cluster to create a snapshot for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_cluster_snapshot#db_cluster_identifier RdsClusterSnapshot#db_cluster_identifier}
	DbClusterIdentifier *string `field:"required" json:"dbClusterIdentifier" yaml:"dbClusterIdentifier"`
	// The identifier for the DB cluster snapshot.
	//
	// Must contain from 1 to 63 letters, numbers, or hyphens. First character must be a letter. Can't end with a hyphen or contain two consecutive hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_cluster_snapshot#db_cluster_snapshot_identifier RdsClusterSnapshot#db_cluster_snapshot_identifier}
	DbClusterSnapshotIdentifier *string `field:"required" json:"dbClusterSnapshotIdentifier" yaml:"dbClusterSnapshotIdentifier"`
	// The tags to be assigned to the DB cluster snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/rds_cluster_snapshot#tags RdsClusterSnapshot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

