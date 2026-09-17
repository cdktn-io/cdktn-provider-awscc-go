// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentDeploymentConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_inference_component#auto_rollback_configuration SagemakerInferenceComponent#auto_rollback_configuration}.
	AutoRollbackConfiguration *SagemakerInferenceComponentDeploymentConfigAutoRollbackConfiguration `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// The rolling update policy for the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_inference_component#rolling_update_policy SagemakerInferenceComponent#rolling_update_policy}
	RollingUpdatePolicy *SagemakerInferenceComponentDeploymentConfigRollingUpdatePolicy `field:"optional" json:"rollingUpdatePolicy" yaml:"rollingUpdatePolicy"`
}

