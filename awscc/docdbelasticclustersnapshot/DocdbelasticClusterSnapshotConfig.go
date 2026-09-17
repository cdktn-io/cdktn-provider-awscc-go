// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package docdbelasticclustersnapshot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DocdbelasticClusterSnapshotConfig struct {
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
	// The ARN of the elastic cluster of which to create a snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/docdbelastic_cluster_snapshot#cluster_arn DocdbelasticClusterSnapshot#cluster_arn}
	ClusterArn *string `field:"required" json:"clusterArn" yaml:"clusterArn"`
	// The name of the elastic cluster snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/docdbelastic_cluster_snapshot#snapshot_name DocdbelasticClusterSnapshot#snapshot_name}
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// An array of key-value pairs to apply to the elastic cluster snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/docdbelastic_cluster_snapshot#tags DocdbelasticClusterSnapshot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

