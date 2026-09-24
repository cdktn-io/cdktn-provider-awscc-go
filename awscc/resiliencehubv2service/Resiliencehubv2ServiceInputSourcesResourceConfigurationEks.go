// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resiliencehubv2service


type Resiliencehubv2ServiceInputSourcesResourceConfigurationEks struct {
	// ARN of the EKS cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_service#cluster_arn Resiliencehubv2Service#cluster_arn}
	ClusterArn *string `field:"optional" json:"clusterArn" yaml:"clusterArn"`
	// Kubernetes label selector that scopes discovery to matching objects in the specified namespaces.
	//
	// An object must satisfy both MatchLabels and MatchExpressions. Specify at least one of them; a selector carrying neither is treated as though no selector were supplied, and all supported objects in the specified namespaces are discovered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_service#label_selector Resiliencehubv2Service#label_selector}
	LabelSelector *Resiliencehubv2ServiceInputSourcesResourceConfigurationEksLabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// EKS namespaces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/resiliencehubv2_service#namespaces Resiliencehubv2Service#namespaces}
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
}

