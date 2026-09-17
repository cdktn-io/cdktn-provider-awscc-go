// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerprocessingjob


type SagemakerProcessingJobProcessingOutputConfig struct {
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the processing job output.
	//
	// KmsKeyId can be an ID of a KMS key, ARN of a KMS key, or alias of a KMS key. The KmsKeyId is applied to all outputs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_processing_job#kms_key_id SagemakerProcessingJob#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// An array of outputs configuring the data to upload from the processing container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_processing_job#outputs SagemakerProcessingJob#outputs}
	Outputs interface{} `field:"optional" json:"outputs" yaml:"outputs"`
}

