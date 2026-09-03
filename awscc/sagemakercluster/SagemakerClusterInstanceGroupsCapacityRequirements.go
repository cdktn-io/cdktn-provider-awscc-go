// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsCapacityRequirements struct {
	// Options for OnDemand capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#on_demand SagemakerCluster#on_demand}
	OnDemand *string `field:"optional" json:"onDemand" yaml:"onDemand"`
	// Options for Spot capacity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#spot SagemakerCluster#spot}
	Spot *string `field:"optional" json:"spot" yaml:"spot"`
}

