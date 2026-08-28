// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscluster


type PcsClusterScheduler struct {
	// The software AWS PCS uses to manage cluster scaling and job scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcs_cluster#type PcsCluster#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// The version of the specified scheduling software that AWS PCS uses to manage cluster scaling and job scheduling.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/pcs_cluster#version PcsCluster#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

