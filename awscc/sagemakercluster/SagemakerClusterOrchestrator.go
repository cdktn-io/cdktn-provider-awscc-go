// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterOrchestrator struct {
	// Specifies parameter(s) related to EKS as orchestrator, e.g. the EKS cluster nodes will attach to,.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#eks SagemakerCluster#eks}
	Eks *SagemakerClusterOrchestratorEks `field:"optional" json:"eks" yaml:"eks"`
	// Specifies parameter(s) related to Slurm as orchestrator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/sagemaker_cluster#slurm SagemakerCluster#slurm}
	Slurm *SagemakerClusterOrchestratorSlurm `field:"optional" json:"slurm" yaml:"slurm"`
}

