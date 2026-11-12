// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListener struct {
	// The IP address on which the destination endpoint listens for incoming connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#destination_address MediaconnectDestinationEndpoint#destination_address}
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// The port on which the destination endpoint listens for incoming connections. The valid range is 0 to 65531.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#destination_ingest_port MediaconnectDestinationEndpoint#destination_ingest_port}
	DestinationIngestPort *float64 `field:"optional" json:"destinationIngestPort" yaml:"destinationIngestPort"`
	// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#encryption_passphrase MediaconnectDestinationEndpoint#encryption_passphrase}
	EncryptionPassphrase *MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphrase `field:"optional" json:"encryptionPassphrase" yaml:"encryptionPassphrase"`
	// The minimum latency, in milliseconds, for the SRT connection. The valid range is 10 to 10000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#minimum_latency_milliseconds MediaconnectDestinationEndpoint#minimum_latency_milliseconds}
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The stream ID that identifies the content for the SRT connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#stream_id MediaconnectDestinationEndpoint#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

