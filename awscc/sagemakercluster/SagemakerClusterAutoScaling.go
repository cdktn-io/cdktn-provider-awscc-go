// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterAutoScaling struct {
	// The type of auto-scaler to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster#auto_scaler_type SagemakerCluster#auto_scaler_type}
	AutoScalerType *string `field:"optional" json:"autoScalerType" yaml:"autoScalerType"`
	// The auto-scaling mode for the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_cluster#mode SagemakerCluster#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

