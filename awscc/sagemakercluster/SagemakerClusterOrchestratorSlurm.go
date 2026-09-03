// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterOrchestratorSlurm struct {
	// The strategy for managing Slurm configuration on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#slurm_config_strategy SagemakerCluster#slurm_config_strategy}
	SlurmConfigStrategy *string `field:"optional" json:"slurmConfigStrategy" yaml:"slurmConfigStrategy"`
}

