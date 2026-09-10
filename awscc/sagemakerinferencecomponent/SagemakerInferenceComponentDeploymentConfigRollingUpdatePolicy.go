// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicy struct {
	// Capacity size configuration for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#maximum_batch_size SagemakerInferenceComponent#maximum_batch_size}
	MaximumBatchSize *SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyMaximumBatchSize `field:"optional" json:"maximumBatchSize" yaml:"maximumBatchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#maximum_execution_timeout_in_seconds SagemakerInferenceComponent#maximum_execution_timeout_in_seconds}.
	MaximumExecutionTimeoutInSeconds *float64 `field:"optional" json:"maximumExecutionTimeoutInSeconds" yaml:"maximumExecutionTimeoutInSeconds"`
	// Capacity size configuration for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#rollback_maximum_batch_size SagemakerInferenceComponent#rollback_maximum_batch_size}
	RollbackMaximumBatchSize *SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicyRollbackMaximumBatchSize `field:"optional" json:"rollbackMaximumBatchSize" yaml:"rollbackMaximumBatchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#wait_interval_in_seconds SagemakerInferenceComponent#wait_interval_in_seconds}.
	WaitIntervalInSeconds *float64 `field:"optional" json:"waitIntervalInSeconds" yaml:"waitIntervalInSeconds"`
}

