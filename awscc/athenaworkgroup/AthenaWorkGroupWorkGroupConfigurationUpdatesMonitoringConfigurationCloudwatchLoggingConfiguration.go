// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationUpdatesMonitoringConfigurationCloudwatchLoggingConfiguration struct {
	// Enables CloudWatch logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#enabled AthenaWorkGroup#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The name of the log group in Amazon CloudWatch Logs where you want to publish your logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#log_group AthenaWorkGroup#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// Prefix for the CloudWatch log stream name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#log_stream_name_prefix AthenaWorkGroup#log_stream_name_prefix}
	LogStreamNamePrefix *string `field:"optional" json:"logStreamNamePrefix" yaml:"logStreamNamePrefix"`
	// The types of logs that you want to publish to CloudWatch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/athena_work_group#log_types AthenaWorkGroup#log_types}
	LogTypes interface{} `field:"optional" json:"logTypes" yaml:"logTypes"`
}

