// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouteroutput


type MediaconnectRouterOutputConfigurationMediaLiveInput struct {
	// The encryption configuration that defines how content is encrypted during transit between MediaConnect Router and MediaLive.
	//
	// This configuration determines whether encryption keys are automatically managed by the service or manually managed through Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_output#destination_transit_encryption MediaconnectRouterOutput#destination_transit_encryption}
	DestinationTransitEncryption *MediaconnectRouterOutputConfigurationMediaLiveInputDestinationTransitEncryption `field:"optional" json:"destinationTransitEncryption" yaml:"destinationTransitEncryption"`
	// The ARN of the MediaLive input to connect to this router output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_output#media_live_input_arn MediaconnectRouterOutput#media_live_input_arn}
	MediaLiveInputArn *string `field:"optional" json:"mediaLiveInputArn" yaml:"mediaLiveInputArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_output#media_live_pipeline_id MediaconnectRouterOutput#media_live_pipeline_id}.
	MediaLivePipelineId *string `field:"optional" json:"mediaLivePipelineId" yaml:"mediaLivePipelineId"`
}

