// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionModelBiasJobOutputConfigMonitoringOutputs struct {
	// Information about where and how to store the results of a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_model_bias_job_definition#s3_output SagemakerModelBiasJobDefinition#s3_output}
	S3Output *SagemakerModelBiasJobDefinitionModelBiasJobOutputConfigMonitoringOutputsS3Output `field:"required" json:"s3Output" yaml:"s3Output"`
}

