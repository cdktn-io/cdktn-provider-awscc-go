// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakerinferencecomponent


type SagemakerInferenceComponentSpecificationComputeResourceRequirements struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_inference_component#max_memory_required_in_mb SagemakerInferenceComponent#max_memory_required_in_mb}.
	MaxMemoryRequiredInMb *float64 `field:"optional" json:"maxMemoryRequiredInMb" yaml:"maxMemoryRequiredInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_inference_component#min_memory_required_in_mb SagemakerInferenceComponent#min_memory_required_in_mb}.
	MinMemoryRequiredInMb *float64 `field:"optional" json:"minMemoryRequiredInMb" yaml:"minMemoryRequiredInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_inference_component#number_of_accelerator_devices_required SagemakerInferenceComponent#number_of_accelerator_devices_required}.
	NumberOfAcceleratorDevicesRequired *float64 `field:"optional" json:"numberOfAcceleratorDevicesRequired" yaml:"numberOfAcceleratorDevicesRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_inference_component#number_of_cpu_cores_required SagemakerInferenceComponent#number_of_cpu_cores_required}.
	NumberOfCpuCoresRequired *float64 `field:"optional" json:"numberOfCpuCoresRequired" yaml:"numberOfCpuCoresRequired"`
}

