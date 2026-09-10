// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ec2verifiedaccessinstance


type Ec2VerifiedAccessInstanceLoggingConfigurations struct {
	// Sends Verified Access logs to CloudWatch Logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_instance#cloudwatch_logs Ec2VerifiedAccessInstance#cloudwatch_logs}
	CloudwatchLogs *Ec2VerifiedAccessInstanceLoggingConfigurationsCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Include claims from trust providers in Verified Access logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_instance#include_trust_context Ec2VerifiedAccessInstance#include_trust_context}
	IncludeTrustContext interface{} `field:"optional" json:"includeTrustContext" yaml:"includeTrustContext"`
	// Sends Verified Access logs to Kinesis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_instance#kinesis_data_firehose Ec2VerifiedAccessInstance#kinesis_data_firehose}
	KinesisDataFirehose *Ec2VerifiedAccessInstanceLoggingConfigurationsKinesisDataFirehose `field:"optional" json:"kinesisDataFirehose" yaml:"kinesisDataFirehose"`
	// Select log version for Verified Access logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_instance#log_version Ec2VerifiedAccessInstance#log_version}
	LogVersion *string `field:"optional" json:"logVersion" yaml:"logVersion"`
	// Sends Verified Access logs to Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/ec2_verified_access_instance#s3 Ec2VerifiedAccessInstance#s3}
	S3 *Ec2VerifiedAccessInstanceLoggingConfigurationsS3 `field:"optional" json:"s3" yaml:"s3"`
}

