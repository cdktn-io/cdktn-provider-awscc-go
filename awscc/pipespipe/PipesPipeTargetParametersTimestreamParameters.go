// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersTimestreamParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pipes_pipe#dimension_mappings PipesPipe#dimension_mappings}.
	DimensionMappings interface{} `field:"optional" json:"dimensionMappings" yaml:"dimensionMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pipes_pipe#epoch_time_unit PipesPipe#epoch_time_unit}.
	EpochTimeUnit *string `field:"optional" json:"epochTimeUnit" yaml:"epochTimeUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pipes_pipe#multi_measure_mappings PipesPipe#multi_measure_mappings}.
	MultiMeasureMappings interface{} `field:"optional" json:"multiMeasureMappings" yaml:"multiMeasureMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pipes_pipe#single_measure_mappings PipesPipe#single_measure_mappings}.
	SingleMeasureMappings interface{} `field:"optional" json:"singleMeasureMappings" yaml:"singleMeasureMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pipes_pipe#time_field_type PipesPipe#time_field_type}.
	TimeFieldType *string `field:"optional" json:"timeFieldType" yaml:"timeFieldType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pipes_pipe#timestamp_format PipesPipe#timestamp_format}.
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pipes_pipe#time_value PipesPipe#time_value}.
	TimeValue *string `field:"optional" json:"timeValue" yaml:"timeValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pipes_pipe#version_value PipesPipe#version_value}.
	VersionValue *string `field:"optional" json:"versionValue" yaml:"versionValue"`
}

