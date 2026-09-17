// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsKubernetesConfig struct {
	// A map of Kubernetes labels to apply to cluster nodes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#labels SagemakerCluster#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// A list of Kubernetes taints to apply to cluster nodes. Maximum of 50 taints.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/sagemaker_cluster#taints SagemakerCluster#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
}

