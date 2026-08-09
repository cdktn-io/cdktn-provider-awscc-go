// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagent


type BedrockAgentActionGroupsApiSchemaS3 struct {
	// A bucket in S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#s3_bucket_name BedrockAgent#s3_bucket_name}
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// A object key in S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/bedrock_agent#s3_object_key BedrockAgent#s3_object_key}
	S3ObjectKey *string `field:"optional" json:"s3ObjectKey" yaml:"s3ObjectKey"`
}

