// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup


type PcsComputeNodeGroupSlurmConfigurationSlurmCustomSettings struct {
	// AWS PCS supports configuration of the following Slurm parameters for compute node groups: Weight and RealMemory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pcs_compute_node_group#parameter_name PcsComputeNodeGroup#parameter_name}
	ParameterName *string `field:"optional" json:"parameterName" yaml:"parameterName"`
	// The value for the configured Slurm setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/pcs_compute_node_group#parameter_value PcsComputeNodeGroup#parameter_value}
	ParameterValue *string `field:"optional" json:"parameterValue" yaml:"parameterValue"`
}

