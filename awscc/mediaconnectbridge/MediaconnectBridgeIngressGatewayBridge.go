// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridge


type MediaconnectBridgeIngressGatewayBridge struct {
	// The maximum expected bitrate of the ingress bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge#max_bitrate MediaconnectBridge#max_bitrate}
	MaxBitrate *float64 `field:"optional" json:"maxBitrate" yaml:"maxBitrate"`
	// The maximum number of outputs on the ingress bridge.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge#max_outputs MediaconnectBridge#max_outputs}
	MaxOutputs *float64 `field:"optional" json:"maxOutputs" yaml:"maxOutputs"`
}

