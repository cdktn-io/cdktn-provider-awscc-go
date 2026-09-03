// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chimechannelflow


type ChimeChannelFlowProcessors struct {
	// A processor's metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_channel_flow#configuration ChimeChannelFlow#configuration}
	Configuration *ChimeChannelFlowProcessorsConfiguration `field:"required" json:"configuration" yaml:"configuration"`
	// The sequence in which processors run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_channel_flow#execution_order ChimeChannelFlow#execution_order}
	ExecutionOrder *float64 `field:"required" json:"executionOrder" yaml:"executionOrder"`
	// Determines whether to continue or stop processing when communication with a processor fails.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_channel_flow#fallback_action ChimeChannelFlow#fallback_action}
	FallbackAction *string `field:"required" json:"fallbackAction" yaml:"fallbackAction"`
	// The name of the processor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/chime_channel_flow#name ChimeChannelFlow#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

