// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterOutpostConfigControlPlanePlacement struct {
	// The name of the placement group for the Kubernetes control plane instances.
	//
	// This setting can't be changed after cluster creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#group_name EksCluster#group_name}
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Optional parameter to specify the placement group spread level for control plane instances.
	//
	// If not provided, EKS will deploy control plane instances without a placement group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#spread_level EksCluster#spread_level}
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
}

