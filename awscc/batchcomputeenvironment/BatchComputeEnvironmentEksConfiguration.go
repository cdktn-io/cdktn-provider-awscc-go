// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package batchcomputeenvironment


type BatchComputeEnvironmentEksConfiguration struct {
	// The EKS access entry configuration for the compute environment.
	//
	// Controls whether AWS Batch manages the EKS access entry for the compute environment's service role, or inherits it from the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/batch_compute_environment#access_entry BatchComputeEnvironment#access_entry}
	AccessEntry *BatchComputeEnvironmentEksConfigurationAccessEntry `field:"optional" json:"accessEntry" yaml:"accessEntry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/batch_compute_environment#eks_cluster_arn BatchComputeEnvironment#eks_cluster_arn}.
	EksClusterArn *string `field:"optional" json:"eksClusterArn" yaml:"eksClusterArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/batch_compute_environment#kubernetes_namespace BatchComputeEnvironment#kubernetes_namespace}.
	KubernetesNamespace *string `field:"optional" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
}

