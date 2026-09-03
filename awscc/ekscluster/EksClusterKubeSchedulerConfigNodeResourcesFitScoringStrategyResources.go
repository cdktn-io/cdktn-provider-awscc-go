// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeSchedulerConfigNodeResourcesFitScoringStrategyResources struct {
	// The name of the resource (for example, cpu or memory).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_cluster#name EksCluster#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The weight assigned to the resource for scoring. Must be between 1 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/eks_cluster#weight EksCluster#weight}
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

