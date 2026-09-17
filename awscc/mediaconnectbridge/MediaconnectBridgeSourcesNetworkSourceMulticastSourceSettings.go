// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridge


type MediaconnectBridgeSourcesNetworkSourceMulticastSourceSettings struct {
	// The IP address of the source for source-specific multicast (SSM).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge#multicast_source_ip MediaconnectBridge#multicast_source_ip}
	MulticastSourceIp *string `field:"optional" json:"multicastSourceIp" yaml:"multicastSourceIp"`
}

