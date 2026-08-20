// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterControlPlaneScalingConfig struct {
	// The scaling tier for the provisioned control plane.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/eks_cluster#tier EksCluster#tier}
	Tier *string `field:"optional" json:"tier" yaml:"tier"`
}

