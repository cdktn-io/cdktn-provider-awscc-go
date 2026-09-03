// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscomputenodegroup


type PcsComputeNodeGroupCustomLaunchTemplate struct {
	// The version of the EC2 launch template to use to provision instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/pcs_compute_node_group#version PcsComputeNodeGroup#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// The ID of the EC2 launch template to use to provision instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/pcs_compute_node_group#template_id PcsComputeNodeGroup#template_id}
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
}

