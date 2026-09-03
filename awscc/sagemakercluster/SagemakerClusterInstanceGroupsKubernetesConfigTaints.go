// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsKubernetesConfigTaints struct {
	// The effect of the taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#effect SagemakerCluster#effect}
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// The key of the taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#key SagemakerCluster#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of the taint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#value SagemakerCluster#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

