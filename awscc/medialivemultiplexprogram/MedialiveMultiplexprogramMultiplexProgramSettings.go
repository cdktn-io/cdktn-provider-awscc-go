// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package medialivemultiplexprogram


type MedialiveMultiplexprogramMultiplexProgramSettings struct {
	// Indicates which pipeline is preferred by the multiplex for program ingest.
	//
	// If set to \"PIPELINE_0\" or \"PIPELINE_1\" and an unhealthy ingest causes the multiplex to switch to the non-preferred pipeline,
	// it will switch back once that ingest is healthy again. If set to \"CURRENTLY_ACTIVE\",
	// it will not switch back to the other pipeline based on it recovering to a healthy state,
	// it will only switch if the active pipeline becomes unhealthy.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/medialive_multiplexprogram#preferred_channel_pipeline MedialiveMultiplexprogram#preferred_channel_pipeline}
	PreferredChannelPipeline *string `field:"optional" json:"preferredChannelPipeline" yaml:"preferredChannelPipeline"`
	// Unique program number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/medialive_multiplexprogram#program_number MedialiveMultiplexprogram#program_number}
	ProgramNumber *float64 `field:"optional" json:"programNumber" yaml:"programNumber"`
	// Transport stream service descriptor configuration for the Multiplex program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/medialive_multiplexprogram#service_descriptor MedialiveMultiplexprogram#service_descriptor}
	ServiceDescriptor *MedialiveMultiplexprogramMultiplexProgramSettingsServiceDescriptor `field:"optional" json:"serviceDescriptor" yaml:"serviceDescriptor"`
	// Program video settings configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/medialive_multiplexprogram#video_settings MedialiveMultiplexprogram#video_settings}
	VideoSettings *MedialiveMultiplexprogramMultiplexProgramSettingsVideoSettings `field:"optional" json:"videoSettings" yaml:"videoSettings"`
}

