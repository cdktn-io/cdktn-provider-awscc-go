// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowSourceFailoverConfig struct {
	// The type of failover you choose for this flow.
	//
	// MERGE combines the source streams into a single stream, allowing graceful recovery from any single-source loss. FAILOVER allows switching between different streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#failover_mode MediaconnectFlow#failover_mode}
	FailoverMode *string `field:"optional" json:"failoverMode" yaml:"failoverMode"`
	// Search window time to look for dash-7 packets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#recovery_window MediaconnectFlow#recovery_window}
	RecoveryWindow *float64 `field:"optional" json:"recoveryWindow" yaml:"recoveryWindow"`
	// The priority you want to assign to a source.
	//
	// You can have a primary stream and a backup stream or two equally prioritized streams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#source_priority MediaconnectFlow#source_priority}
	SourcePriority *MediaconnectFlowSourceFailoverConfigSourcePriority `field:"optional" json:"sourcePriority" yaml:"sourcePriority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/mediaconnect_flow#state MediaconnectFlow#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

