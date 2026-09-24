// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationTransformations struct {
	// When the service applies the transformation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#step_to_apply BedrockDataSource#step_to_apply}
	StepToApply *string `field:"optional" json:"stepToApply" yaml:"stepToApply"`
	// A Lambda function that processes documents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#transformation_function BedrockDataSource#transformation_function}
	TransformationFunction *BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationTransformationsTransformationFunction `field:"optional" json:"transformationFunction" yaml:"transformationFunction"`
}

