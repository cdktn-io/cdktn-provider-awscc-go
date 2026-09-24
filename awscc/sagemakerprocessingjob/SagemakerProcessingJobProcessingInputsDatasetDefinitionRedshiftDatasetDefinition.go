// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingInputsDatasetDefinitionRedshiftDatasetDefinition struct {
	// The Redshift cluster Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#cluster_id SagemakerProcessingJob#cluster_id}
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
	// The IAM role attached to your Redshift cluster that Amazon SageMaker uses to generate datasets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#cluster_role_arn SagemakerProcessingJob#cluster_role_arn}
	ClusterRoleArn *string `field:"optional" json:"clusterRoleArn" yaml:"clusterRoleArn"`
	// The name of the Redshift database used in Redshift query execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#database SagemakerProcessingJob#database}
	Database *string `field:"optional" json:"database" yaml:"database"`
	// The database user name used in Redshift query execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#db_user SagemakerProcessingJob#db_user}
	DbUser *string `field:"optional" json:"dbUser" yaml:"dbUser"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt data from a Redshift execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#kms_key_id SagemakerProcessingJob#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The compression used for Redshift query results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#output_compression SagemakerProcessingJob#output_compression}
	OutputCompression *string `field:"optional" json:"outputCompression" yaml:"outputCompression"`
	// The data storage format for Redshift query results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#output_format SagemakerProcessingJob#output_format}
	OutputFormat *string `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// The location in Amazon S3 where the Redshift query results are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#output_s3_uri SagemakerProcessingJob#output_s3_uri}
	OutputS3Uri *string `field:"optional" json:"outputS3Uri" yaml:"outputS3Uri"`
	// The SQL query statements to be executed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_processing_job#query_string SagemakerProcessingJob#query_string}
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
}

