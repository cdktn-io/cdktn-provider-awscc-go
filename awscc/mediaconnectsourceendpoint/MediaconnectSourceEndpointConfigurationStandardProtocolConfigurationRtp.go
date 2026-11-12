// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationRtp struct {
	// The port on the destination that receives the content. The valid range is 3000 to 30000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#destination_ingest_port MediaconnectSourceEndpoint#destination_ingest_port}
	DestinationIngestPort *float64 `field:"optional" json:"destinationIngestPort" yaml:"destinationIngestPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#forward_error_correction_state MediaconnectSourceEndpoint#forward_error_correction_state}.
	ForwardErrorCorrectionState *string `field:"optional" json:"forwardErrorCorrectionState" yaml:"forwardErrorCorrectionState"`
}

