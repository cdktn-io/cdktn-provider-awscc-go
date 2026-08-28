// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinition struct {
	// The name of the data catalog used in Athena query execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#catalog SagemakerProcessingJob#catalog}
	Catalog *string `field:"optional" json:"catalog" yaml:"catalog"`
	// The name of the database used in the Athena query execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#database SagemakerProcessingJob#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data generated from an Athena query execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#kms_key_id SagemakerProcessingJob#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The compression used for Athena query results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#output_compression SagemakerProcessingJob#output_compression}
	OutputCompression *string `field:"optional" json:"outputCompression" yaml:"outputCompression"`
	// The data storage format for Athena query results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#output_format SagemakerProcessingJob#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The location in Amazon S3 where Athena query results are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#output_s3_uri SagemakerProcessingJob#output_s3_uri}
	OutputS3Uri *string `field:"optional" json:"outputS3Uri" yaml:"outputS3Uri"`
	// The SQL query statements, to be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#query_string SagemakerProcessingJob#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// The name of the workgroup in which the Athena query is being started.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#work_group SagemakerProcessingJob#work_group}
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

