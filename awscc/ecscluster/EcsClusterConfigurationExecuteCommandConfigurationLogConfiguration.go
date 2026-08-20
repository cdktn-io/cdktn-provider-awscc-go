// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ecscluster


type EcsClusterConfigurationExecuteCommandConfigurationLogConfiguration struct {
	// Determines whether to use encryption on the CloudWatch logs. If not specified, encryption will be off.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_cluster#cloudwatch_encryption_enabled EcsCluster#cloudwatch_encryption_enabled}
	CloudwatchEncryptionEnabled interface{} `field:"optional" json:"cloudwatchEncryptionEnabled" yaml:"cloudwatchEncryptionEnabled"`
	// The name of the CloudWatch log group to send logs to.
	//
	// The CloudWatch log group must already be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_cluster#cloudwatch_log_group_name EcsCluster#cloudwatch_log_group_name}
	CloudwatchLogGroupName *string `field:"optional" json:"cloudwatchLogGroupName" yaml:"cloudwatchLogGroupName"`
	// The name of the S3 bucket to send logs to.   The S3 bucket must already be created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_cluster#s3_bucket_name EcsCluster#s3_bucket_name}
	S3BucketName *string `field:"optional" json:"s3BucketName" yaml:"s3BucketName"`
	// Determines whether to use encryption on the S3 logs. If not specified, encryption is not used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_cluster#s3_encryption_enabled EcsCluster#s3_encryption_enabled}
	S3EncryptionEnabled interface{} `field:"optional" json:"s3EncryptionEnabled" yaml:"s3EncryptionEnabled"`
	// An optional folder in the S3 bucket to place logs in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/ecs_cluster#s3_key_prefix EcsCluster#s3_key_prefix}
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

