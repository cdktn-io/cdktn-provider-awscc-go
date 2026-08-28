// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionModelBiasJobOutputConfig struct {
	// Monitoring outputs for monitoring jobs. This is where the output of the periodic monitoring jobs is uploaded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_bias_job_definition#monitoring_outputs SagemakerModelBiasJobDefinition#monitoring_outputs}
	MonitoringOutputs interface{} `field:"required" json:"monitoringOutputs" yaml:"monitoringOutputs"`
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the model artifacts at rest using Amazon S3 server-side encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/sagemaker_model_bias_job_definition#kms_key_id SagemakerModelBiasJobDefinition#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

