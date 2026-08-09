// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcsqueue


type PcsQueueComputeNodeGroupConfigurations struct {
	// The compute node group ID for the compute node group configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_queue#compute_node_group_id PcsQueue#compute_node_group_id}
	ComputeNodeGroupId *string `field:"optional" json:"computeNodeGroupId" yaml:"computeNodeGroupId"`
}

