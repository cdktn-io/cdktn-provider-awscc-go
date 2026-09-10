// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelqualityjobdefinition


type SagemakerModelQualityJobDefinitionModelQualityJobInput struct {
	// Ground truth input provided in S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_quality_job_definition#ground_truth_s3_input SagemakerModelQualityJobDefinition#ground_truth_s3_input}
	GroundTruthS3Input *SagemakerModelQualityJobDefinitionModelQualityJobInputGroundTruthS3Input `field:"required" json:"groundTruthS3Input" yaml:"groundTruthS3Input"`
	// The batch transform input for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_quality_job_definition#batch_transform_input SagemakerModelQualityJobDefinition#batch_transform_input}
	BatchTransformInput *SagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInput `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// The endpoint for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_model_quality_job_definition#endpoint_input SagemakerModelQualityJobDefinition#endpoint_input}
	EndpointInput *SagemakerModelQualityJobDefinitionModelQualityJobInputEndpointInput `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

