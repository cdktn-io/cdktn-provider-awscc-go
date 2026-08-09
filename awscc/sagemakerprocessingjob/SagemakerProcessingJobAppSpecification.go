// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobAppSpecification struct {
	// The container image to be run by the processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_processing_job#image_uri SagemakerProcessingJob#image_uri}
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The arguments for a container used to run a processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_processing_job#container_arguments SagemakerProcessingJob#container_arguments}
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// The entrypoint for a container used to run a processing job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_processing_job#container_entrypoint SagemakerProcessingJob#container_entrypoint}
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
}

