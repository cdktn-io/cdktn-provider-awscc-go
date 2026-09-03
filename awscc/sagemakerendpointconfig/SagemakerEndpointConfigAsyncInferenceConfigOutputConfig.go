// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigAsyncInferenceConfigOutputConfig struct {
	// The AWS Key Management Service (AWS KMS) key that Amazon SageMaker uses to encrypt the asynchronous inference output in Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#kms_key_id SagemakerEndpointConfigA#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Specifies the configuration for notifications of inference results for asynchronous inference.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#notification_config SagemakerEndpointConfigA#notification_config}
	NotificationConfig *SagemakerEndpointConfigAsyncInferenceConfigOutputConfigNotificationConfig `field:"optional" json:"notificationConfig" yaml:"notificationConfig"`
	// The Amazon S3 location to upload failure inference responses to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#s3_failure_path SagemakerEndpointConfigA#s3_failure_path}
	S3FailurePath *string `field:"optional" json:"s3FailurePath" yaml:"s3FailurePath"`
	// The Amazon S3 location to upload inference responses to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_endpoint_config#s3_output_path SagemakerEndpointConfigA#s3_output_path}
	S3OutputPath *string `field:"optional" json:"s3OutputPath" yaml:"s3OutputPath"`
}

