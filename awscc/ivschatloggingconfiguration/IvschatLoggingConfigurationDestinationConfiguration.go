// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ivschatloggingconfiguration


type IvschatLoggingConfigurationDestinationConfiguration struct {
	// CloudWatch destination configuration for IVS Chat logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivschat_logging_configuration#cloudwatch_logs IvschatLoggingConfiguration#cloudwatch_logs}
	CloudwatchLogs *IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Kinesis Firehose destination configuration for IVS Chat logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivschat_logging_configuration#firehose IvschatLoggingConfiguration#firehose}
	Firehose *IvschatLoggingConfigurationDestinationConfigurationFirehose `field:"optional" json:"firehose" yaml:"firehose"`
	// S3 destination configuration for IVS Chat logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/ivschat_logging_configuration#s3 IvschatLoggingConfiguration#s3}
	S3 *IvschatLoggingConfigurationDestinationConfigurationS3 `field:"optional" json:"s3" yaml:"s3"`
}

