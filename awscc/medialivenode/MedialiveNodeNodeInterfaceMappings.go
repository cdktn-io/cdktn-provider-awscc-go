// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivenode


type MedialiveNodeNodeInterfaceMappings struct {
	// The logical name for this interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/medialive_node#logical_interface_name MedialiveNode#logical_interface_name}
	LogicalInterfaceName *string `field:"optional" json:"logicalInterfaceName" yaml:"logicalInterfaceName"`
	// The network interface mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/medialive_node#network_interface_mode MedialiveNode#network_interface_mode}
	NetworkInterfaceMode *string `field:"optional" json:"networkInterfaceMode" yaml:"networkInterfaceMode"`
	// The physical interface name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/medialive_node#physical_interface_name MedialiveNode#physical_interface_name}
	PhysicalInterfaceName *string `field:"optional" json:"physicalInterfaceName" yaml:"physicalInterfaceName"`
}

