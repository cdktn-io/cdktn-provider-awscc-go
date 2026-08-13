// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationFailoverProtocolConfigurationsSrtListener struct {
	// Contains the configuration settings for decrypting SRT streams, including the encryption key details and decryption parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#decryption_configuration MediaconnectRouterInput#decryption_configuration}
	DecryptionConfiguration *MediaconnectRouterInputConfigurationFailoverProtocolConfigurationsSrtListenerDecryptionConfiguration `field:"optional" json:"decryptionConfiguration" yaml:"decryptionConfiguration"`
	// The minimum latency in milliseconds for the SRT protocol in listener mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#minimum_latency_milliseconds MediaconnectRouterInput#minimum_latency_milliseconds}
	MinimumLatencyMilliseconds *float64 `field:"optional" json:"minimumLatencyMilliseconds" yaml:"minimumLatencyMilliseconds"`
	// The port number for the SRT protocol in listener mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_router_input#port MediaconnectRouterInput#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

