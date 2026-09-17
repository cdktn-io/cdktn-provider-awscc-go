// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationMediaLiveChannel struct {
	// The ARN of the MediaLive channel to connect to this router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_router_input#media_live_channel_arn MediaconnectRouterInput#media_live_channel_arn}
	MediaLiveChannelArn *string `field:"optional" json:"mediaLiveChannelArn" yaml:"mediaLiveChannelArn"`
	// The name of the MediaLive channel output to connect to this router input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_router_input#media_live_channel_output_name MediaconnectRouterInput#media_live_channel_output_name}
	MediaLiveChannelOutputName *string `field:"optional" json:"mediaLiveChannelOutputName" yaml:"mediaLiveChannelOutputName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_router_input#media_live_pipeline_id MediaconnectRouterInput#media_live_pipeline_id}.
	MediaLivePipelineId *string `field:"optional" json:"mediaLivePipelineId" yaml:"mediaLivePipelineId"`
	// The encryption configuration that defines how content is encrypted during transit between MediaConnect Router and MediaLive.
	//
	// This configuration determines whether encryption keys are automatically managed by the service or manually managed through Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/mediaconnect_router_input#source_transit_decryption MediaconnectRouterInput#source_transit_decryption}
	SourceTransitDecryption *MediaconnectRouterInputConfigurationMediaLiveChannelSourceTransitDecryption `field:"optional" json:"sourceTransitDecryption" yaml:"sourceTransitDecryption"`
}

