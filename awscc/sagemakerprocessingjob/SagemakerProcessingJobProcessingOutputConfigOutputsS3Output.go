// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingOutputConfigOutputsS3Output struct {
	// The local path of a directory where you want Amazon SageMaker to upload its contents to Amazon S3.
	//
	// LocalPath is an absolute path to a directory containing output files. This directory will be created by the platform and exist when your container's entrypoint is invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_processing_job#local_path SagemakerProcessingJob#local_path}
	LocalPath *string `field:"optional" json:"localPath" yaml:"localPath"`
	// Whether to upload the results of the processing job continuously or after the job completes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_processing_job#s3_upload_mode SagemakerProcessingJob#s3_upload_mode}
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
	// A URI that identifies the Amazon S3 bucket where you want Amazon SageMaker to save the results of a processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_processing_job#s3_uri SagemakerProcessingJob#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

