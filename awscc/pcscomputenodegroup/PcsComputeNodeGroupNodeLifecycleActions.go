// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup


type PcsComputeNodeGroupNodeLifecycleActions struct {
	// Controls whether lifecycle scripts are downloaded once at first boot (CACHE_ONCE) or re-downloaded on every reboot (REFRESH_ON_REBOOT).
	//
	// Defaults to CACHE_ONCE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pcs_compute_node_group#script_caching_policy PcsComputeNodeGroup#script_caching_policy}
	ScriptCachingPolicy *string `field:"optional" json:"scriptCachingPolicy" yaml:"scriptCachingPolicy"`
	// The ordered scripts to run at each compute node lifecycle stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/pcs_compute_node_group#stages PcsComputeNodeGroup#stages}
	Stages *PcsComputeNodeGroupNodeLifecycleActionsStages `field:"optional" json:"stages" yaml:"stages"`
}

