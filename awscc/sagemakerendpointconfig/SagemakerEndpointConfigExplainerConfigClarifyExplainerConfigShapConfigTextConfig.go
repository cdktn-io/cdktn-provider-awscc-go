// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfig


type SagemakerEndpointConfigExplainerConfigClarifyExplainerConfigShapConfigTextConfig struct {
	// The unit of granularity for the analysis of text features.
	//
	// For example, if the unit is 'token', then each token (like a word in English) of the text is treated as a feature. SHAP values are computed for each unit/feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint_config#granularity SagemakerEndpointConfigA#granularity}
	Granularity *string `field:"optional" json:"granularity" yaml:"granularity"`
	// Specifies the language of the text features in ISO 639-1 or ISO 639-3 code of a supported language.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_endpoint_config#language SagemakerEndpointConfigA#language}
	Language *string `field:"optional" json:"language" yaml:"language"`
}

