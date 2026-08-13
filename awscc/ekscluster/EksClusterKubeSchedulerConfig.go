// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ekscluster


type EksClusterKubeSchedulerConfig struct {
	// The NodeResourcesFit plugin configuration for the Kubernetes scheduler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/eks_cluster#node_resources_fit EksCluster#node_resources_fit}
	NodeResourcesFit *EksClusterKubeSchedulerConfigNodeResourcesFit `field:"optional" json:"nodeResourcesFit" yaml:"nodeResourcesFit"`
}

