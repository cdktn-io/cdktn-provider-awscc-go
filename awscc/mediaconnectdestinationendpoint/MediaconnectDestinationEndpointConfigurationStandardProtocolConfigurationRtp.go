// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationRtp struct {
	// The IP address on which the destination endpoint receives content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#destination_address MediaconnectDestinationEndpoint#destination_address}
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// The port on which the destination endpoint receives content. The valid range is 1024 to 65531.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#destination_ingest_port MediaconnectDestinationEndpoint#destination_ingest_port}
	DestinationIngestPort *float64 `field:"optional" json:"destinationIngestPort" yaml:"destinationIngestPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#forward_error_correction_state MediaconnectDestinationEndpoint#forward_error_correction_state}.
	ForwardErrorCorrectionState *string `field:"optional" json:"forwardErrorCorrectionState" yaml:"forwardErrorCorrectionState"`
}

