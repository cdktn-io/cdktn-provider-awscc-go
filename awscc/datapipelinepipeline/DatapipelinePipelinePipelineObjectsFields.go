// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datapipelinepipeline


type DatapipelinePipelinePipelineObjectsFields struct {
	// Specifies the name of a field for a particular object.
	//
	// To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datapipeline_pipeline#key DatapipelinePipeline#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A field value that you specify as an identifier of another object in the same pipeline definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datapipeline_pipeline#ref_value DatapipelinePipeline#ref_value}
	RefValue *string `field:"optional" json:"refValue" yaml:"refValue"`
	// A field value that you specify as a string.
	//
	// To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/datapipeline_pipeline#string_value DatapipelinePipeline#string_value}
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

