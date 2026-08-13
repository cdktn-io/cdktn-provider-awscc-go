// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeSchedulerConfigNodeResourcesFitScoringStrategy struct {
	// The resource weights used for scoring nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/eks_cluster#resources EksCluster#resources}
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
	// The scoring strategy type (LeastAllocated or MostAllocated).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/eks_cluster#type EksCluster#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

