// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationTransformationsTransformationFunction struct {
	// A Lambda function that processes documents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/bedrock_data_source#transformation_lambda_configuration BedrockDataSource#transformation_lambda_configuration}
	TransformationLambdaConfiguration *BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationTransformationsTransformationFunctionTransformationLambdaConfiguration `field:"optional" json:"transformationLambdaConfiguration" yaml:"transformationLambdaConfiguration"`
}

