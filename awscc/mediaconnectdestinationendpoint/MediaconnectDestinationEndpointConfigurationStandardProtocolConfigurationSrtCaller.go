// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectdestinationendpoint


type MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtCaller struct {
	// The configuration for the Secrets Manager secret that holds the key or passphrase used to encrypt an endpoint connection.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#encryption_passphrase MediaconnectDestinationEndpoint#encryption_passphrase}
	EncryptionPassphrase *MediaconnectDestinationEndpointConfigurationStandardProtocolConfigurationSrtCallerEncryptionPassphrase `field:"optional" json:"encryptionPassphrase" yaml:"encryptionPassphrase"`
	// The minimum latency, in milliseconds, for the SRT connection. The valid range is 10 to 10000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#minimum_latency_milliseconds MediaconnectDestinationEndpoint#minimum_latency_milliseconds}
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The port on which the source listens for the connection. The valid range is 3000 to 30000.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_destination_endpoint#source_listener_port MediaconnectDestinationEndpoint#source_listener_port}
	SourceListenerPort *float64 `field:"optional" json:"sourceListenerPort" yaml:"sourceListenerPort"`
}

