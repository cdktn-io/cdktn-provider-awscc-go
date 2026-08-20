// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterOrchestratorEks struct {
	// The ARN of the EKS cluster, such as arn:aws:eks:us-west-2:123456789012:cluster/my-eks-cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/sagemaker_cluster#cluster_arn SagemakerCluster#cluster_arn}
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
}

