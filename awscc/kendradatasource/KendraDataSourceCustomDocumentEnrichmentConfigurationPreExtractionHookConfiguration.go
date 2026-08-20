// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kendradatasource


type KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#invocation_condition KendraDataSource#invocation_condition}.
	InvocationCondition *KendraDataSourceCustomDocumentEnrichmentConfigurationPreExtractionHookConfigurationInvocationCondition `field:"optional" json:"invocationCondition" yaml:"invocationCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#lambda_arn KendraDataSource#lambda_arn}.
	LambdaArn *string `field:"optional" json:"lambdaArn" yaml:"lambdaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kendra_data_source#s3_bucket KendraDataSource#s3_bucket}.
	S3Bucket *string `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

