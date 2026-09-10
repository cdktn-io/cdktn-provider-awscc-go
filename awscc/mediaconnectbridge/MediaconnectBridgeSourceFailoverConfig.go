// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectbridge


type MediaconnectBridgeSourceFailoverConfig struct {
	// The type of failover you choose for this flow. FAILOVER allows switching between different streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#failover_mode MediaconnectBridge#failover_mode}
	FailoverMode *string `field:"optional" json:"failoverMode" yaml:"failoverMode"`
	// The priority you want to assign to a source.
	//
	// You can have a primary stream and a backup stream or two equally prioritized streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#source_priority MediaconnectBridge#source_priority}
	SourcePriority *MediaconnectBridgeSourceFailoverConfigSourcePriority `field:"optional" json:"sourcePriority" yaml:"sourcePriority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/mediaconnect_bridge#state MediaconnectBridge#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

