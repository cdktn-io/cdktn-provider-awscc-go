// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingInputs struct {
	// When True, input operations such as data download are managed natively by the processing job application.
	//
	// When False (default), input operations are managed by Amazon SageMaker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#app_managed SagemakerProcessingJob#app_managed}
	AppManaged interface{} `field:"optional" json:"appManaged" yaml:"appManaged"`
	// Configuration for Dataset Definition inputs. The Dataset Definition input must specify exactly one of either `AthenaDatasetDefinition` or `RedshiftDatasetDefinition` types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#dataset_definition SagemakerProcessingJob#dataset_definition}
	DatasetDefinition *SagemakerProcessingJobProcessingInputsDatasetDefinition `field:"optional" json:"datasetDefinition" yaml:"datasetDefinition"`
	// The name for the processing job input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#input_name SagemakerProcessingJob#input_name}
	InputName *string `field:"optional" json:"inputName" yaml:"inputName"`
	// Configuration for downloading input data from Amazon S3 into the processing container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/sagemaker_processing_job#s3_input SagemakerProcessingJob#s3_input}
	S3Input *SagemakerProcessingJobProcessingInputsS3Input `field:"optional" json:"s3Input" yaml:"s3Input"`
}

