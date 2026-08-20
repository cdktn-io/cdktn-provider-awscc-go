// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pcsqueue


type PcsQueueSlurmConfiguration struct {
	// Custom Slurm parameters that directly map to Slurm configuration settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/pcs_queue#slurm_custom_settings PcsQueue#slurm_custom_settings}
	SlurmCustomSettings interface{} `field:"optional" json:"slurmCustomSettings" yaml:"slurmCustomSettings"`
}

