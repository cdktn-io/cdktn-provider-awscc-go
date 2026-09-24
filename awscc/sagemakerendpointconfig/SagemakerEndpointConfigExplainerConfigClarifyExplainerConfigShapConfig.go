// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfig struct {
	// The number of samples to be used for analysis by the Kernal SHAP algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#number_of_samples SagemakerEndpointConfigA#number_of_samples}
	NumberOfSamples *float64 `field:"optional" json:"numberOfSamples" yaml:"numberOfSamples"`
	// The starting value used to initialize the random number generator in the explainer.
	//
	// Provide a value for this parameter to obtain a deterministic SHAP result.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#seed SagemakerEndpointConfigA#seed}
	Seed *float64 `field:"optional" json:"seed" yaml:"seed"`
	// The configuration for the SHAP baseline of the Kernal SHAP algorithm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#shap_baseline_config SagemakerEndpointConfigA#shap_baseline_config}
	ShapBaselineConfig *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigShapBaselineConfig `field:"optional" json:"shapBaselineConfig" yaml:"shapBaselineConfig"`
	// A parameter that indicates if text features are treated as text and explanations are provided for individual units of text.
	//
	// Required for natural language processing (NLP) explainability only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#text_config SagemakerEndpointConfigA#text_config}
	TextConfig *SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigTextConfig `field:"optional" json:"textConfig" yaml:"textConfig"`
	// A Boolean toggle to indicate if you want to use the logit function (true) or log-odds units (false) for model predictions.
	//
	// Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_endpoint_config#use_logit SagemakerEndpointConfigA#use_logit}
	UseLogit interface{} `field:"optional" json:"useLogit" yaml:"useLogit"`
}

