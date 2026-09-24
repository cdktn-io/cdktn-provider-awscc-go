// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferenceexperiment


type SagemakerInferenceExperimentModelVariantsInfrastructureConfig struct {
	// The type of the inference experiment that you want to run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_experiment#infrastructure_type SagemakerInferenceExperiment#infrastructure_type}
	InfrastructureType *string `field:"required" json:"infrastructureType" yaml:"infrastructureType"`
	// The infrastructure configuration for deploying the model to a real-time inference endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_experiment#real_time_inference_config SagemakerInferenceExperiment#real_time_inference_config}
	RealTimeInferenceConfig *SagemakerInferenceExperimentModelVariantsInfrastructureConfigRealTimeInferenceConfig `field:"required" json:"realTimeInferenceConfig" yaml:"realTimeInferenceConfig"`
}

