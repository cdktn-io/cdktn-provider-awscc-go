// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionModelBiasJobInput struct {
	// Ground truth input provided in S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_bias_job_definition#ground_truth_s3_input SagemakerModelBiasJobDefinition#ground_truth_s3_input}
	GroundTruthS3Input *SagemakerModelBiasJobDefinitionModelBiasJobInputGroundTruthS3Input `field:"required" json:"groundTruthS3Input" yaml:"groundTruthS3Input"`
	// The batch transform input for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_bias_job_definition#batch_transform_input SagemakerModelBiasJobDefinition#batch_transform_input}
	BatchTransformInput *SagemakerModelBiasJobDefinitionModelBiasJobInputBatchTransformInput `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// The endpoint for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_model_bias_job_definition#endpoint_input SagemakerModelBiasJobDefinition#endpoint_input}
	EndpointInput *SagemakerModelBiasJobDefinitionModelBiasJobInputEndpointInput `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

