// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferenceexperiment


type SagemakerInferenceExperimentShadowModeConfigShadowModelVariants struct {
	// The percentage of inference requests that Amazon SageMaker replicates from the production variant to the shadow variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_inference_experiment#sampling_percentage SagemakerInferenceExperiment#sampling_percentage}
	SamplingPercentage *float64 `field:"optional" json:"samplingPercentage" yaml:"samplingPercentage"`
	// The name of the shadow variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_inference_experiment#shadow_model_variant_name SagemakerInferenceExperiment#shadow_model_variant_name}
	ShadowModelVariantName *string `field:"optional" json:"shadowModelVariantName" yaml:"shadowModelVariantName"`
}

