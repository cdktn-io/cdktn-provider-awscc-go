// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowNdiConfigNdiDiscoveryServers struct {
	// The unique network address of the NDI discovery server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#discovery_server_address MediaconnectFlow#discovery_server_address}
	DiscoveryServerAddress *string `field:"optional" json:"discoveryServerAddress" yaml:"discoveryServerAddress"`
	// The port for the NDI discovery server. Defaults to 5959 if a custom port isn't specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#discovery_server_port MediaconnectFlow#discovery_server_port}
	DiscoveryServerPort *float64 `field:"optional" json:"discoveryServerPort" yaml:"discoveryServerPort"`
	// The identifier for the Virtual Private Cloud (VPC) network interface used by the flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_flow#vpc_interface_adapter MediaconnectFlow#vpc_interface_adapter}
	VpcInterfaceAdapter *string `field:"optional" json:"vpcInterfaceAdapter" yaml:"vpcInterfaceAdapter"`
}

