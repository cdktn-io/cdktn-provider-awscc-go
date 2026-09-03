// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package omicsruncache

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type OmicsRunCacheConfig struct {
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
	// The default cache behavior for runs using this cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_run_cache#cache_behavior OmicsRunCache#cache_behavior}
	CacheBehavior *string `field:"optional" json:"cacheBehavior" yaml:"cacheBehavior"`
	// The AWS account ID of the expected owner of the S3 bucket for the run cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_run_cache#cache_bucket_owner_id OmicsRunCache#cache_bucket_owner_id}
	CacheBucketOwnerId *string `field:"optional" json:"cacheBucketOwnerId" yaml:"cacheBucketOwnerId"`
	// The S3 location for storing the cached task outputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_run_cache#cache_s3_location OmicsRunCache#cache_s3_location}
	CacheS3Location *string `field:"optional" json:"cacheS3Location" yaml:"cacheS3Location"`
	// A description of the run cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_run_cache#description OmicsRunCache#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// A name for the run cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_run_cache#name OmicsRunCache#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Tags for the run cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/omics_run_cache#tags OmicsRunCache#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

