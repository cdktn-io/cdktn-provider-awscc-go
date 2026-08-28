// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingInputsDatasetDefinition struct {
	// Configuration for Athena Dataset Definition input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#athena_dataset_definition SagemakerProcessingJob#athena_dataset_definition}
	AthenaDatasetDefinition *SagemakerProcessingJobProcessingInputsDatasetDefinitionAthenaDatasetDefinition `field:"optional" json:"athenaDatasetDefinition" yaml:"athenaDatasetDefinition"`
	// Whether the generated dataset is FullyReplicated or ShardedByS3Key (default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#data_distribution_type SagemakerProcessingJob#data_distribution_type}
	DataDistributionType *string `field:"optional" json:"dataDistributionType" yaml:"dataDistributionType"`
	// Whether to use File or Pipe input mode.
	//
	// In File (default) mode, Amazon SageMaker copies the data from the input source onto the local Amazon Elastic Block Store (Amazon EBS) volumes before starting your training algorithm. This is the most commonly used input mode. In Pipe mode, Amazon SageMaker streams input data from the source directly to your algorithm without using the EBS volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#input_mode SagemakerProcessingJob#input_mode}
	InputMode *string `field:"optional" json:"inputMode" yaml:"inputMode"`
	// The local path where you want Amazon SageMaker to download the Dataset Definition inputs to run a processing job.
	//
	// LocalPath is an absolute path to the input data. This is a required parameter when AppManaged is False (default).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#local_path SagemakerProcessingJob#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Configuration for Redshift Dataset Definition input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_processing_job#redshift_dataset_definition SagemakerProcessingJob#redshift_dataset_definition}
	RedshiftDatasetDefinition *SagemakerProcessingJobProcessingInputsDatasetDefinitionRedshiftDatasetDefinition `field:"optional" json:"redshiftDatasetDefinition" yaml:"redshiftDatasetDefinition"`
}

