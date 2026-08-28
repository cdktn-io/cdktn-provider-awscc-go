// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersTimestreamParametersDimensionMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#dimension_name PipesPipe#dimension_name}.
	DimensionName *string `field:"optional" json:"dimensionName" yaml:"dimensionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#dimension_value PipesPipe#dimension_value}.
	DimensionValue *string `field:"optional" json:"dimensionValue" yaml:"dimensionValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pipes_pipe#dimension_value_type PipesPipe#dimension_value_type}.
	DimensionValueType *string `field:"optional" json:"dimensionValueType" yaml:"dimensionValueType"`
}

