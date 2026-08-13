// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lambdalayerversion


type LambdaLayerVersionContent struct {
	// The Amazon S3 bucket of the layer archive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lambda_layer_version#s3_bucket LambdaLayerVersion#s3_bucket}
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// The Amazon S3 key of the layer archive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lambda_layer_version#s3_key LambdaLayerVersion#s3_key}
	S3Key *string `field:"required" json:"s3Key" yaml:"s3Key"`
	// Specifies whether Lambda should copy the deployment package to its internal storage (COPY) or reference it directly from your S3 bucket (REFERENCE).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lambda_layer_version#s3_object_storage_mode LambdaLayerVersion#s3_object_storage_mode}
	S3ObjectStorageMode *string `field:"optional" json:"s3ObjectStorageMode" yaml:"s3ObjectStorageMode"`
	// For versioned objects, the version of the layer archive object to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/lambda_layer_version#s3_object_version LambdaLayerVersion#s3_object_version}
	S3ObjectVersion *string `field:"optional" json:"s3ObjectVersion" yaml:"s3ObjectVersion"`
}

