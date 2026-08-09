// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterTieredStorageConfig struct {
	// The percentage of instance memory to allocate for tiered storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_cluster#instance_memory_allocation_percentage SagemakerCluster#instance_memory_allocation_percentage}
	InstanceMemoryAllocationPercentage *float64 `field:"optional" json:"instanceMemoryAllocationPercentage" yaml:"instanceMemoryAllocationPercentage"`
	// The mode of tiered storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/sagemaker_cluster#mode SagemakerCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

