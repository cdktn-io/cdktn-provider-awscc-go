// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercomputequota


type SagemakerComputeQuotaComputeQuotaConfigComputeQuotaResourcesAcceleratorPartition struct {
	// The number of MIG partitions of the specified Type allocated to the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#count SagemakerComputeQuota#count}
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// MIG profile type. Service rejects MIG profiles that the chosen InstanceType does not support.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.107.0/docs/resources/sagemaker_compute_quota#type SagemakerComputeQuota#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

