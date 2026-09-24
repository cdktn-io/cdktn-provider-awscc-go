// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridgeoutput


type MediaconnectBridgeOutputNetworkOutput struct {
	// The network output IP Address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge_output#ip_address MediaconnectBridgeOutput#ip_address}
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// The network output's gateway network name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge_output#network_name MediaconnectBridgeOutput#network_name}
	NetworkName *string `field:"required" json:"networkName" yaml:"networkName"`
	// The network output port.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge_output#port MediaconnectBridgeOutput#port}
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// The network output protocol.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge_output#protocol MediaconnectBridgeOutput#protocol}
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The network output TTL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_bridge_output#ttl MediaconnectBridgeOutput#ttl}
	Ttl *float64 `field:"required" json:"ttl" yaml:"ttl"`
}

