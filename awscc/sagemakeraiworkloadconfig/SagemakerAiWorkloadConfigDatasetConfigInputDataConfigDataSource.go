// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeraiworkloadconfig


type SagemakerAiWorkloadConfigDatasetConfigInputDataConfigDataSource struct {
	// The Amazon S3 data source configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_ai_workload_config#s3_data_source SagemakerAiWorkloadConfig#s3_data_source}
	S3DataSource *SagemakerAiWorkloadConfigDatasetConfigInputDataConfigDataSourceS3DataSource `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}

