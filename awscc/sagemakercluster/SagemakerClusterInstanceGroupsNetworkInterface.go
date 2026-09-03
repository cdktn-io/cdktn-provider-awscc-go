// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakercluster


type SagemakerClusterInstanceGroupsNetworkInterface struct {
	// The type of network interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/sagemaker_cluster#interface_type SagemakerCluster#interface_type}
	InterfaceType *string `field:"optional" json:"interfaceType" yaml:"interfaceType"`
}

