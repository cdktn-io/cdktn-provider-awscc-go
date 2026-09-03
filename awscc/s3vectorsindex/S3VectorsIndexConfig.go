// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package s3vectorsindex

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type S3VectorsIndexConfig struct {
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
	// The data type of the vectors to be inserted into the vector index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3vectors_index#data_type S3VectorsIndex#data_type}
	DataType *string `field:"required" json:"dataType" yaml:"dataType"`
	// The dimensions of the vectors to be inserted into the vector index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3vectors_index#dimension S3VectorsIndex#dimension}
	Dimension *float64 `field:"required" json:"dimension" yaml:"dimension"`
	// The distance metric to be used for similarity search.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3vectors_index#distance_metric S3VectorsIndex#distance_metric}
	DistanceMetric *string `field:"required" json:"distanceMetric" yaml:"distanceMetric"`
	// The encryption configuration for the index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3vectors_index#encryption_configuration S3VectorsIndex#encryption_configuration}
	EncryptionConfiguration *S3VectorsIndexEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// The name of the vector index to create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3vectors_index#index_name S3VectorsIndex#index_name}
	IndexName *string `field:"optional" json:"indexName" yaml:"indexName"`
	// The metadata configuration for the vector index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3vectors_index#metadata_configuration S3VectorsIndex#metadata_configuration}
	MetadataConfiguration *S3VectorsIndexMetadataConfiguration `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// User tags (key-value pairs) to associate with the index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3vectors_index#tags S3VectorsIndex#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The Amazon Resource Name (ARN) of the vector bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3vectors_index#vector_bucket_arn S3VectorsIndex#vector_bucket_arn}
	VectorBucketArn *string `field:"optional" json:"vectorBucketArn" yaml:"vectorBucketArn"`
	// The name of the vector bucket that contains the vector index.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/s3vectors_index#vector_bucket_name S3VectorsIndex#vector_bucket_name}
	VectorBucketName *string `field:"optional" json:"vectorBucketName" yaml:"vectorBucketName"`
}

