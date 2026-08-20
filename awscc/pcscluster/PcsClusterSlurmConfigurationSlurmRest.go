// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcscluster


type PcsClusterSlurmConfigurationSlurmRest struct {
	// The default value is `NONE`. A value of `STANDARD` means that Slurm Rest is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcs_cluster#mode PcsCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

