// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingOutputConfigOutputs struct {
	// When True, output operations such as data upload are managed natively by the processing job application.
	//
	// When False (default), output operations are managed by Amazon SageMaker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#app_managed SagemakerProcessingJob#app_managed}
	AppManaged interface{} `field:"optional" json:"appManaged" yaml:"appManaged"`
	// Configuration for processing job outputs in Amazon SageMaker Feature Store.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#feature_store_output SagemakerProcessingJob#feature_store_output}
	FeatureStoreOutput *SagemakerProcessingJobProcessingOutputConfigOutputsFeatureStoreOutput `field:"optional" json:"featureStoreOutput" yaml:"featureStoreOutput"`
	// The name for the processing job output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#output_name SagemakerProcessingJob#output_name}
	OutputName *string `field:"optional" json:"outputName" yaml:"outputName"`
	// Configuration for uploading output data to Amazon S3 from the processing container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_processing_job#s3_output SagemakerProcessingJob#s3_output}
	S3Output *SagemakerProcessingJobProcessingOutputConfigOutputsS3Output `field:"optional" json:"s3Output" yaml:"s3Output"`
}

