// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceInputSourcesResourceConfigurationEks struct {
	// ARN of the EKS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resiliencehubv2_service#cluster_arn Resiliencehubv2Service#cluster_arn}
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// EKS namespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/resiliencehubv2_service#namespaces Resiliencehubv2Service#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

