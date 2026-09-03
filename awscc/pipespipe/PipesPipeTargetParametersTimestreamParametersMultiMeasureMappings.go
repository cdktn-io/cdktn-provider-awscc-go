// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pipespipe


type PipesPipeTargetParametersTimestreamParametersMultiMeasureMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/pipes_pipe#multi_measure_attribute_mappings PipesPipe#multi_measure_attribute_mappings}.
	MultiMeasureAttributeMappings interface{} `field:"optional" json:"multiMeasureAttributeMappings" yaml:"multiMeasureAttributeMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/pipes_pipe#multi_measure_name PipesPipe#multi_measure_name}.
	MultiMeasureName *string `field:"optional" json:"multiMeasureName" yaml:"multiMeasureName"`
}

