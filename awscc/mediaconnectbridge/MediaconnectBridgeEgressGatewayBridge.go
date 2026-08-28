// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridge


type MediaconnectBridgeEgressGatewayBridge struct {
	// The maximum expected bitrate of the egress bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_bridge#max_bitrate MediaconnectBridge#max_bitrate}
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
}

