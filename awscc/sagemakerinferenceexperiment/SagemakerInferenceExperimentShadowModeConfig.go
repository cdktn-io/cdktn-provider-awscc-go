// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferenceexperiment


type SagemakerInferenceExperimentShadowModeConfig struct {
	// List of shadow variant configurations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_experiment#shadow_model_variants SagemakerInferenceExperiment#shadow_model_variants}
	ShadowModelVariants interface{} `field:"optional" json:"shadowModelVariants" yaml:"shadowModelVariants"`
	// The name of the production variant, which takes all the inference requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_experiment#source_model_variant_name SagemakerInferenceExperiment#source_model_variant_name}
	SourceModelVariantName *string `field:"optional" json:"sourceModelVariantName" yaml:"sourceModelVariantName"`
}

