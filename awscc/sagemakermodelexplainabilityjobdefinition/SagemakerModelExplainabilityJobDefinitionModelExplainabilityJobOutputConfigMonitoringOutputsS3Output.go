// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobOutputConfigMonitoringOutputsS3Output struct {
	// The local path to the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	//
	// LocalPath is an absolute path for the output data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_explainability_job_definition#local_path SagemakerModelExplainabilityJobDefinition#local_path}
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_explainability_job_definition#s3_uri SagemakerModelExplainabilityJobDefinition#s3_uri}
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
	// Whether to upload the results of the monitoring job continuously or after the job completes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_explainability_job_definition#s3_upload_mode SagemakerModelExplainabilityJobDefinition#s3_upload_mode}
	S3UploadMode *string `field:"optional" json:"s3UploadMode" yaml:"s3UploadMode"`
}

