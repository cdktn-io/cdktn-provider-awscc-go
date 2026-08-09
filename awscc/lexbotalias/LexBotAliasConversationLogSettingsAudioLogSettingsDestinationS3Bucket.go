// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package lexbotalias


type LexBotAliasConversationLogSettingsAudioLogSettingsDestinationS3Bucket struct {
	// The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot_alias#kms_key_arn LexBotAlias#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// The Amazon S3 key of the deployment package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot_alias#log_prefix LexBotAlias#log_prefix}
	LogPrefix *string `field:"optional" json:"logPrefix" yaml:"logPrefix"`
	// The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/lex_bot_alias#s3_bucket_arn LexBotAlias#s3_bucket_arn}
	S3BucketArn *string `field:"optional" json:"s3BucketArn" yaml:"s3BucketArn"`
}

