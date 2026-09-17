// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelqualityjobdefinition


type SagemakerModelQualityJobDefinitionModelQualityAppSpecification struct {
	// The container image to be run by the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#image_uri SagemakerModelQualityJobDefinition#image_uri}
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The status of the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#problem_type SagemakerModelQualityJobDefinition#problem_type}
	ProblemType *string `field:"required" json:"problemType" yaml:"problemType"`
	// An array of arguments for the container used to run the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#container_arguments SagemakerModelQualityJobDefinition#container_arguments}
	ContainerArguments *[]*string `field:"optional" json:"containerArguments" yaml:"containerArguments"`
	// Specifies the entrypoint for a container used to run the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#container_entrypoint SagemakerModelQualityJobDefinition#container_entrypoint}
	ContainerEntrypoint *[]*string `field:"optional" json:"containerEntrypoint" yaml:"containerEntrypoint"`
	// Sets the environment variables in the Docker container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#environment SagemakerModelQualityJobDefinition#environment}
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// An Amazon S3 URI to a script that is called after analysis has been performed.
	//
	// Applicable only for the built-in (first party) containers.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#post_analytics_processor_source_uri SagemakerModelQualityJobDefinition#post_analytics_processor_source_uri}
	PostAnalyticsProcessorSourceUri *string `field:"optional" json:"postAnalyticsProcessorSourceUri" yaml:"postAnalyticsProcessorSourceUri"`
	// An Amazon S3 URI to a script that is called per row prior to running analysis.
	//
	// It can base64 decode the payload and convert it into a flatted json so that the built-in container can use the converted data. Applicable only for the built-in (first party) containers
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_quality_job_definition#record_preprocessor_source_uri SagemakerModelQualityJobDefinition#record_preprocessor_source_uri}
	RecordPreprocessorSourceUri *string `field:"optional" json:"recordPreprocessorSourceUri" yaml:"recordPreprocessorSourceUri"`
}

