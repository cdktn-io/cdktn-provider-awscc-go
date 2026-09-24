// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferenceexperiment


type SagemakerInferenceExperimentModelVariants struct {
	// The configuration for the infrastructure that the model will be deployed to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_experiment#infrastructure_config SagemakerInferenceExperiment#infrastructure_config}
	InfrastructureConfig *SagemakerInferenceExperimentModelVariantsInfrastructureConfig `field:"required" json:"infrastructureConfig" yaml:"infrastructureConfig"`
	// The name of the Amazon SageMaker Model entity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_experiment#model_name SagemakerInferenceExperiment#model_name}
	ModelName *string `field:"required" json:"modelName" yaml:"modelName"`
	// The name of the variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_experiment#variant_name SagemakerInferenceExperiment#variant_name}
	VariantName *string `field:"required" json:"variantName" yaml:"variantName"`
}

