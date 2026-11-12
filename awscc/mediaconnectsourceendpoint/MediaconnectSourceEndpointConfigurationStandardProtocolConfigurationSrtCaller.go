// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCaller struct {
	// The port on which the destination listens for the connection. The valid range is 3000 to 30000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#destination_ingest_port MediaconnectSourceEndpoint#destination_ingest_port}
	DestinationIngestPort *float64 `field:"optional" json:"destinationIngestPort" yaml:"destinationIngestPort"`
	// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#encryption_passphrase MediaconnectSourceEndpoint#encryption_passphrase}
	EncryptionPassphrase *MediaconnectSourceEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphrase `field:"optional" json:"encryptionPassphrase" yaml:"encryptionPassphrase"`
	// The minimum latency, in milliseconds, for the SRT connection. The valid range is 10 to 10000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#minimum_latency_milliseconds MediaconnectSourceEndpoint#minimum_latency_milliseconds}
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
}

