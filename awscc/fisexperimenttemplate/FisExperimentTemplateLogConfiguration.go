// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package fisexperimenttemplate


type FisExperimentTemplateLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#cloudwatch_logs_configuration FisExperimentTemplate#cloudwatch_logs_configuration}.
	CloudwatchLogsConfiguration *FisExperimentTemplateLogConfigurationCloudwatchLogsConfiguration `field:"optional" json:"cloudwatchLogsConfiguration" yaml:"cloudwatchLogsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#log_schema_version FisExperimentTemplate#log_schema_version}.
	LogSchemaVersion *float64 `field:"optional" json:"logSchemaVersion" yaml:"logSchemaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/fis_experiment_template#s3_configuration FisExperimentTemplate#s3_configuration}.
	S3Configuration *FisExperimentTemplateLogConfigurationS3Configuration `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

