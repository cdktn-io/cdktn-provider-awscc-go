// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridge


type MediaconnectBridgeOutputs struct {
	// The output of the bridge. A network output is delivered to your premises.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_bridge#network_output MediaconnectBridge#network_output}
	NetworkOutput *MediaconnectBridgeOutputsNetworkOutput `field:"optional" json:"networkOutput" yaml:"networkOutput"`
}

