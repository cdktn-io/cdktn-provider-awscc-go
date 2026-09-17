// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridgesource


type MediaconnectBridgeSourceNetworkSource struct {
	// The network source multicast IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#multicast_ip MediaconnectBridgeSource#multicast_ip}
	MulticastIp *string `field:"optional" json:"multicastIp" yaml:"multicastIp"`
	// The settings related to the multicast source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#multicast_source_settings MediaconnectBridgeSource#multicast_source_settings}
	MulticastSourceSettings *MediaconnectBridgeSourceNetworkSourceMulticastSourceSettings `field:"optional" json:"multicastSourceSettings" yaml:"multicastSourceSettings"`
	// The network source's gateway network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#network_name MediaconnectBridgeSource#network_name}
	NetworkName *string `field:"optional" json:"networkName" yaml:"networkName"`
	// The network source port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#port MediaconnectBridgeSource#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The network source protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge_source#protocol MediaconnectBridgeSource#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

