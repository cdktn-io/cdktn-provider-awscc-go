// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3vectorsvectorbucket

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3VectorsVectorBucketConfig struct {
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
	// The encryption configuration for the vector bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3vectors_vector_bucket#encryption_configuration S3VectorsVectorBucket#encryption_configuration}
	EncryptionConfiguration *S3VectorsVectorBucketEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// User tags (key-value pairs) to associate with the vector bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3vectors_vector_bucket#tags S3VectorsVectorBucket#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The name of the vector bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/s3vectors_vector_bucket#vector_bucket_name S3VectorsVectorBucket#vector_bucket_name}
	VectorBucketName *string `field:"optional" json:"vectorBucketName" yaml:"vectorBucketName"`
}

