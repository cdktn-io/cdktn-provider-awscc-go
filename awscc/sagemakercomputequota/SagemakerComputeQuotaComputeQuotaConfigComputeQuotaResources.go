// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercomputequota


type SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResources struct {
	// Multi-instance GPU (MIG) partition profile and count, for fractional accelerator allocation.
	//
	// Service rejects MIG profiles that the chosen InstanceType does not support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#accelerator_partition SagemakerComputeQuota#accelerator_partition}
	AcceleratorPartition *SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesAcceleratorPartition `field:"optional" json:"acceleratorPartition" yaml:"acceleratorPartition"`
	// The number of accelerators (e.g., GPUs) of this type allocated to the team. SageMaker auto-fills VCpu / MemoryInGiB ratio-proportionally if only this is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#accelerators SagemakerComputeQuota#accelerators}
	Accelerators *float64 `field:"optional" json:"accelerators" yaml:"accelerators"`
	// The number of instances of this type allocated to the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#count SagemakerComputeQuota#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// The cluster instance type the quota applies to (e.g., 'ml.m5.xlarge', 'ml.p5.48xlarge'). Validated by the service against the supported HyperPod instance types.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#instance_type SagemakerComputeQuota#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
	// The memory in GiB allocated to the team.
	//
	// Floating-point. SageMaker auto-fills Accelerators / VCpu ratio-proportionally if only this is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#memory_in_gi_b SagemakerComputeQuota#memory_in_gi_b}
	MemoryInGiB *float64 `field:"optional" json:"memoryInGiB" yaml:"memoryInGiB"`
	// The number of vCPUs allocated to the team.
	//
	// Floating-point. SageMaker auto-fills Accelerators / MemoryInGiB ratio-proportionally if only this is set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#v_cpu SagemakerComputeQuota#v_cpu}
	VCpu *float64 `field:"optional" json:"vCpu" yaml:"vCpu"`
}

