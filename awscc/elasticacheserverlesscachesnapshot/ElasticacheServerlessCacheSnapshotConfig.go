// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package elasticacheserverlesscachesnapshot

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type ElasticacheServerlessCacheSnapshotConfig struct {
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
	// The name of an existing serverless cache. The snapshot is created from this cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_serverless_cache_snapshot#serverless_cache_name ElasticacheServerlessCacheSnapshot#serverless_cache_name}
	ServerlessCacheName *string `field:"required" json:"serverlessCacheName" yaml:"serverlessCacheName"`
	// The name of the serverless cache snapshot.
	//
	// Must be unique for the customer account. This value is stored as a lowercase string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_serverless_cache_snapshot#serverless_cache_snapshot_name ElasticacheServerlessCacheSnapshot#serverless_cache_snapshot_name}
	ServerlessCacheSnapshotName *string `field:"required" json:"serverlessCacheSnapshotName" yaml:"serverlessCacheSnapshotName"`
	// The Amazon Resource Name (ARN) of the AWS KMS key used to encrypt the snapshot.
	//
	// Provide the key ARN: the resource returns the key ARN on read, so supplying a bare key ID or alias for this createOnly property may be reported as drift by CloudFormation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_serverless_cache_snapshot#kms_key_id ElasticacheServerlessCacheSnapshot#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// A list of tags to be added to the serverless cache snapshot resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/elasticache_serverless_cache_snapshot#tags ElasticacheServerlessCacheSnapshot#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

