// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterOutpostConfigEtcdPlacement struct {
	// Optional parameter to specify the placement group spread level for etcd instances.
	//
	// If not provided, EKS will deploy etcd instances without a placement group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/eks_cluster#spread_level EksCluster#spread_level}
	SpreadLevel *string `field:"optional" json:"spreadLevel" yaml:"spreadLevel"`
}

