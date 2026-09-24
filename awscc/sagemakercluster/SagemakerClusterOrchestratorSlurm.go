// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterOrchestratorSlurm struct {
	// External MySQL-compatible accounting database that a Slurm cluster's slurmdbd connects to.
	//
	// Database credentials are supplied out-of-band through the referenced Secrets Manager secret. Supported only with Continuous node provisioning.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster#accounting_database SagemakerCluster#accounting_database}
	AccountingDatabase *SagemakerClusterOrchestratorSlurmAccountingDatabase `field:"optional" json:"accountingDatabase" yaml:"accountingDatabase"`
	// The strategy for managing Slurm configuration on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster#slurm_config_strategy SagemakerCluster#slurm_config_strategy}
	SlurmConfigStrategy *string `field:"optional" json:"slurmConfigStrategy" yaml:"slurmConfigStrategy"`
}

