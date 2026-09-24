// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup


type PcsComputeNodeGroupSlurmConfiguration struct {
	// Additional Slurm gres.conf records for the compute node group. Each item is a map of gres.conf attribute names to values describing one gres.conf record (for example a GPU topology, MIG, MPS, or custom GRES entry). AWS PCS adds the NodeName= prefix and merges these records with the GPU record it derives from the instance type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#gres_custom_settings PcsComputeNodeGroup#gres_custom_settings}
	GresCustomSettings interface{} `field:"optional" json:"gresCustomSettings" yaml:"gresCustomSettings"`
	// The time before an idle node is scaled down.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#scale_down_idle_time_in_seconds PcsComputeNodeGroup#scale_down_idle_time_in_seconds}
	ScaleDownIdleTimeInSeconds *float64 `field:"optional" json:"scaleDownIdleTimeInSeconds" yaml:"scaleDownIdleTimeInSeconds"`
	// Additional Slurm-specific configuration that directly maps to Slurm settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/pcs_compute_node_group#slurm_custom_settings PcsComputeNodeGroup#slurm_custom_settings}
	SlurmCustomSettings interface{} `field:"optional" json:"slurmCustomSettings" yaml:"slurmCustomSettings"`
}

