// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup


type PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrapped struct {
	// An ordered list of arguments passed to the script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_compute_node_group#arguments PcsComputeNodeGroup#arguments}
	Arguments *[]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Whether the script runs only on the node's first boot (FIRST_BOOT_ONLY) or on every boot including reboots (EVERY_BOOT).
	//
	// Defaults to FIRST_BOOT_ONLY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_compute_node_group#execution_policy PcsComputeNodeGroup#execution_policy}
	ExecutionPolicy *string `field:"optional" json:"executionPolicy" yaml:"executionPolicy"`
	// A human-readable name that identifies the script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_compute_node_group#name PcsComputeNodeGroup#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The behavior when the script exits with an error. Defaults to TERMINATE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_compute_node_group#on_error PcsComputeNodeGroup#on_error}
	OnError *string `field:"optional" json:"onError" yaml:"onError"`
	// The external location of a lifecycle script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/pcs_compute_node_group#script_source PcsComputeNodeGroup#script_source}
	ScriptSource *PcsComputeNodeGroupNodeLifecycleActionsStagesNodeBootstrappedScriptSource `field:"optional" json:"scriptSource" yaml:"scriptSource"`
}

