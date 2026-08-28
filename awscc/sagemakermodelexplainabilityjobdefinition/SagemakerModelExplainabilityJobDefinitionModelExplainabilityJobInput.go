// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInput struct {
	// The batch transform input for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_explainability_job_definition#batch_transform_input SagemakerModelExplainabilityJobDefinition#batch_transform_input}
	BatchTransformInput *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputBatchTransformInput `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// The endpoint for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_explainability_job_definition#endpoint_input SagemakerModelExplainabilityJobDefinition#endpoint_input}
	EndpointInput *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputEndpointInput `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

