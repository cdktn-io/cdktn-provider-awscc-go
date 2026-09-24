// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigExplainerConfigClarifyExplainerConfig struct {
	// A JMESPath boolean expression used to filter which records to explain. Explanations are activated by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#enable_explanations SagemakerEndpointConfigA#enable_explanations}
	EnableExplanations *string `field:"optional" json:"enableExplanations" yaml:"enableExplanations"`
	// The inference configuration parameter for the model container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#inference_config SagemakerEndpointConfigA#inference_config}
	InferenceConfig *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigInferenceConfig `field:"optional" json:"inferenceConfig" yaml:"inferenceConfig"`
	// The configuration for SHAP analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#shap_config SagemakerEndpointConfigA#shap_config}
	ShapConfig *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfig `field:"optional" json:"shapConfig" yaml:"shapConfig"`
}

