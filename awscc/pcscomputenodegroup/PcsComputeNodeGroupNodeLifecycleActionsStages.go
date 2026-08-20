// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup


type PcsComputeNodeGroupNodeLifecycleActionsStages struct {
	// Scripts to run after the node is bootstrapped, once the PCS configuration phase completes and before slurmd starts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcs_compute_node_group#node_bootstrapped PcsComputeNodeGroup#node_bootstrapped}
	NodeBootstrapped interface{} `field:"optional" json:"nodeBootstrapped" yaml:"nodeBootstrapped"`
	// Scripts to execute when the node becomes ready (every boot).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcs_compute_node_group#node_ready PcsComputeNodeGroup#node_ready}
	NodeReady interface{} `field:"optional" json:"nodeReady" yaml:"nodeReady"`
}

