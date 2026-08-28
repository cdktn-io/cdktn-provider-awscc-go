// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridge


type MediaconnectBridgeSources struct {
	// The source of the bridge. A flow source originates in MediaConnect as an existing cloud flow.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_bridge#flow_source MediaconnectBridge#flow_source}
	FlowSource *MediaconnectBridgeSourcesFlowSource `field:"optional" json:"flowSource" yaml:"flowSource"`
	// The source of the bridge. A network source originates at your premises.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_bridge#network_source MediaconnectBridge#network_source}
	NetworkSource *MediaconnectBridgeSourcesNetworkSource `field:"optional" json:"networkSource" yaml:"networkSource"`
}

