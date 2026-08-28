// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodel


type SagemakerModelContainersModelDataSource struct {
	// Specifies the S3 location of ML model data to deploy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model#s3_data_source SagemakerModel#s3_data_source}
	S3DataSource *SagemakerModelContainersModelDataSourceS3DataSource `field:"optional" json:"s3DataSource" yaml:"s3DataSource"`
}

