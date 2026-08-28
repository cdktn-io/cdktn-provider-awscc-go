// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package imagebuilderinfrastructureconfiguration


type ImagebuilderInfrastructureConfigurationLoggingS3Logs struct {
	// S3BucketName.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_infrastructure_configuration#s3_bucket_name ImagebuilderInfrastructureConfiguration#s3_bucket_name}
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// S3KeyPrefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/imagebuilder_infrastructure_configuration#s3_key_prefix ImagebuilderInfrastructureConfiguration#s3_key_prefix}
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

