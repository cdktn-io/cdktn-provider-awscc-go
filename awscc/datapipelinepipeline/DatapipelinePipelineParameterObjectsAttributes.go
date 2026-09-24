// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DatapipelinePipelineParameterObjectsAttributes struct {
	// The field identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datapipeline_pipeline#key DatapipelinePipeline#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The field value, expressed as a String.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datapipeline_pipeline#string_value DatapipelinePipeline#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

