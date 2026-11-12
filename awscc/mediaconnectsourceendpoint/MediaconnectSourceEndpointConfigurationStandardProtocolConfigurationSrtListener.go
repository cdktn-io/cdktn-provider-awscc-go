// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListener struct {
	// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#encryption_passphrase MediaconnectSourceEndpoint#encryption_passphrase}
	EncryptionPassphrase *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtListenerEncryptionPassphrase `field:"optional" json:"encryptionPassphrase" yaml:"encryptionPassphrase"`
	// The minimum latency, in milliseconds, for the SRT connection. The valid range is 10 to 10000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#minimum_latency_milliseconds MediaconnectSourceEndpoint#minimum_latency_milliseconds}
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The IP address that the source endpoint listens on for incoming connections.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#source_address MediaconnectSourceEndpoint#source_address}
	SourceAddress *string `field:"optional" json:"sourceAddress" yaml:"sourceAddress"`
	// The port that the source endpoint listens on for incoming connections. The valid range is 1024 to 65535.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#source_listener_port MediaconnectSourceEndpoint#source_listener_port}
	SourceListenerPort *float64 `field:"optional" json:"sourceListenerPort" yaml:"sourceListenerPort"`
	// The stream ID that identifies the content for the SRT connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#stream_id MediaconnectSourceEndpoint#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

