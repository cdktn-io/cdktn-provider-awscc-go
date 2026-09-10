// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecification struct {
	// The name of the base inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#base_inference_component_name SagemakerInferenceComponent#base_inference_component_name}
	BaseInferenceComponentName *string `field:"optional" json:"baseInferenceComponentName" yaml:"baseInferenceComponentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#compute_resource_requirements SagemakerInferenceComponent#compute_resource_requirements}.
	ComputeResourceRequirements *SagemakerInferenceComponentSpecificationComputeResourceRequirements `field:"optional" json:"computeResourceRequirements" yaml:"computeResourceRequirements"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#container SagemakerInferenceComponent#container}.
	Container *SagemakerInferenceComponentSpecificationContainer `field:"optional" json:"container" yaml:"container"`
	// Settings that affect how the inference component caches data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#data_cache_config SagemakerInferenceComponent#data_cache_config}
	DataCacheConfig *SagemakerInferenceComponentSpecificationDataCacheConfig `field:"optional" json:"dataCacheConfig" yaml:"dataCacheConfig"`
	// The name of the model to use with the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#model_name SagemakerInferenceComponent#model_name}
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The scheduling configuration that determines how inference component copies are placed across available instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#scheduling_config SagemakerInferenceComponent#scheduling_config}
	SchedulingConfig *SagemakerInferenceComponentSpecificationSchedulingConfig `field:"optional" json:"schedulingConfig" yaml:"schedulingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_inference_component#startup_parameters SagemakerInferenceComponent#startup_parameters}.
	StartupParameters *SagemakerInferenceComponentSpecificationStartupParameters `field:"optional" json:"startupParameters" yaml:"startupParameters"`
}

