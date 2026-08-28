// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationMonitoringConfiguration struct {
	// Configuration settings for delivering logs to Amazon CloudWatch log groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_work_group#cloudwatch_logging_configuration AthenaWorkGroup#cloudwatch_logging_configuration}
	CloudwatchLoggingConfiguration *AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationCloudwatchLoggingConfiguration `field:"optional" json:"cloudwatchLoggingConfiguration" yaml:"cloudwatchLoggingConfiguration"`
	// Configuration settings for managed log persistence.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_work_group#managed_logging_configuration AthenaWorkGroup#managed_logging_configuration}
	ManagedLoggingConfiguration *AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationManagedLoggingConfiguration `field:"optional" json:"managedLoggingConfiguration" yaml:"managedLoggingConfiguration"`
	// Configuration settings for delivering logs to Amazon S3 buckets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/athena_work_group#s3_logging_configuration AthenaWorkGroup#s3_logging_configuration}
	S3LoggingConfiguration *AthenaWorkGroupWorkGroupConfigurationMonitoringConfigurationS3LoggingConfiguration `field:"optional" json:"s3LoggingConfiguration" yaml:"s3LoggingConfiguration"`
}

