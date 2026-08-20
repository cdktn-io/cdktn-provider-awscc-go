// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCaller struct {
	// Contains the configuration settings for decrypting SRT streams, including the encryption key details and decryption parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_input#decryption_configuration MediaconnectRouterInput#decryption_configuration}
	DecryptionConfiguration *MediaconnectRouterInputConfigurationStandardProtocolConfigurationSrtCallerDecryptionConfiguration `field:"optional" json:"decryptionConfiguration" yaml:"decryptionConfiguration"`
	// The minimum latency in milliseconds for the SRT protocol in caller mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_input#minimum_latency_milliseconds MediaconnectRouterInput#minimum_latency_milliseconds}
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The source IP address for the SRT protocol in caller mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_input#source_address MediaconnectRouterInput#source_address}
	SourceAddress *string `field:"optional" json:"sourceAddress" yaml:"sourceAddress"`
	// The source port number for the SRT protocol in caller mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_input#source_port MediaconnectRouterInput#source_port}
	SourcePort *float64 `field:"optional" json:"sourcePort" yaml:"sourcePort"`
	// The stream ID for the SRT protocol in caller mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/mediaconnect_router_input#stream_id MediaconnectRouterInput#stream_id}
	StreamId *string `field:"optional" json:"streamId" yaml:"streamId"`
}

