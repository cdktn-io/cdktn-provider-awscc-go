// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#compute_resource_requirements SagemakerInferenceComponent#compute_resource_requirements}.
	ComputeResourceRequirements *SagemakerInferenceComponentSpecificationsComputeResourceRequirements `field:"optional" json:"computeResourceRequirements" yaml:"computeResourceRequirements"`
	// Container specification for one Specifications entry.
	//
	// Distinct from InferenceComponentContainerSpecification: DescribeInferenceComponent returns no per-entry DeployedImage (VERIFIED in us-west-2), so DeployedImage is intentionally omitted here and this definition can never be aggregated into a plural READ response. The singular InferenceComponentContainerSpecification keeps DeployedImage - the service DOES return it there.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#container SagemakerInferenceComponent#container}
	Container *SagemakerInferenceComponentSpecificationsContainer `field:"optional" json:"container" yaml:"container"`
	// The data caching configuration actually in effect for this instance type, including a value the service chose rather than the template: SageMaker enables caching automatically on instance types with more than 232 GiB of local NVMe storage, whether or not DataCacheConfig was set.
	//
	// Returned by Describe and not settable; set DataCacheConfig instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#current_data_cache_config SagemakerInferenceComponent#current_data_cache_config}
	CurrentDataCacheConfig *SagemakerInferenceComponentSpecificationsCurrentDataCacheConfig `field:"optional" json:"currentDataCacheConfig" yaml:"currentDataCacheConfig"`
	// Settings that affect how the inference component caches data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#data_cache_config SagemakerInferenceComponent#data_cache_config}
	DataCacheConfig *SagemakerInferenceComponentSpecificationsDataCacheConfig `field:"optional" json:"dataCacheConfig" yaml:"dataCacheConfig"`
	// An ML compute instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#instance_type SagemakerInferenceComponent#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The name of the model to use with the inference component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#model_name SagemakerInferenceComponent#model_name}
	ModelName *string `field:"optional" json:"modelName" yaml:"modelName"`
	// The scheduling configuration that determines how inference component copies are placed across available instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#scheduling_config SagemakerInferenceComponent#scheduling_config}
	SchedulingConfig *SagemakerInferenceComponentSpecificationsSchedulingConfig `field:"optional" json:"schedulingConfig" yaml:"schedulingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_inference_component#startup_parameters SagemakerInferenceComponent#startup_parameters}.
	StartupParameters *SagemakerInferenceComponentSpecificationsStartupParameters `field:"optional" json:"startupParameters" yaml:"startupParameters"`
}

