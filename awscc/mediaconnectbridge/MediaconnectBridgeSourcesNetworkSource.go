// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridge


type MediaconnectBridgeSourcesNetworkSource struct {
	// The network source multicast IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge#multicast_ip MediaconnectBridge#multicast_ip}
	MulticastIp *string `field:"optional" json:"multicastIp" yaml:"multicastIp"`
	// The settings related to the multicast source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge#multicast_source_settings MediaconnectBridge#multicast_source_settings}
	MulticastSourceSettings *MediaconnectBridgeSourcesNetworkSourceMulticastSourceSettings `field:"optional" json:"multicastSourceSettings" yaml:"multicastSourceSettings"`
	// The name of the network source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge#name MediaconnectBridge#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The network source's gateway network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge#network_name MediaconnectBridge#network_name}
	NetworkName *string `field:"optional" json:"networkName" yaml:"networkName"`
	// The network source port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge#port MediaconnectBridge#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The network source protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge#protocol MediaconnectBridge#protocol}
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

