// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticachesnapshot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticacheSnapshotConfig struct {
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
	// The name of a snapshot. Must be unique within the customer account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/elasticache_snapshot#snapshot_name ElasticacheSnapshot#snapshot_name}
	SnapshotName *string `field:"required" json:"snapshotName" yaml:"snapshotName"`
	// The identifier of an existing cluster to create a snapshot from. The snapshot is created from this cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/elasticache_snapshot#cache_cluster_id ElasticacheSnapshot#cache_cluster_id}
	CacheClusterId *string `field:"optional" json:"cacheClusterId" yaml:"cacheClusterId"`
	// The ID of the KMS key used to encrypt the snapshot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/elasticache_snapshot#kms_key_id ElasticacheSnapshot#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The identifier of an existing replication group to create a snapshot from.
	//
	// The snapshot is created from this replication group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/elasticache_snapshot#replication_group_id ElasticacheSnapshot#replication_group_id}
	ReplicationGroupId *string `field:"optional" json:"replicationGroupId" yaml:"replicationGroupId"`
	// A list of tags to be added to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/elasticache_snapshot#tags ElasticacheSnapshot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

