// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup


type PcsComputeNodeGroupInstanceConfigs struct {
	// The EC2 instance type that AWS PCS can provision in the compute node group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/pcs_compute_node_group#instance_type PcsComputeNodeGroup#instance_type}
	InstanceType *string `field:"optional" json:"instanceType" yaml:"instanceType"`
}

