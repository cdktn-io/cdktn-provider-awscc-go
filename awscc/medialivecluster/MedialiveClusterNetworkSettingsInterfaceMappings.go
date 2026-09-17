// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivecluster


type MedialiveClusterNetworkSettingsInterfaceMappings struct {
	// logical interface name, unique in the list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/medialive_cluster#logical_interface_name MedialiveCluster#logical_interface_name}
	LogicalInterfaceName *string `field:"optional" json:"logicalInterfaceName" yaml:"logicalInterfaceName"`
	// Network Id to be associated with the logical interface name, can be duplicated in list.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/medialive_cluster#network_id MedialiveCluster#network_id}
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
}

