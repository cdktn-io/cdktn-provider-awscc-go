// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationTransformationsTransformationFunctionTransformationLambdaConfiguration struct {
	// The function's ARN identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/bedrock_data_source#lambda_arn BedrockDataSource#lambda_arn}
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
}

