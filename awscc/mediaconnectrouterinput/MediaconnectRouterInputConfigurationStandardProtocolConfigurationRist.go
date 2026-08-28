// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationStandardProtocolConfigurationRist struct {
	// The port number used for the RIST protocol in the router input configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#port MediaconnectRouterInput#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// The recovery latency in milliseconds for the RIST protocol in the router input configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/mediaconnect_router_input#recovery_latency_milliseconds MediaconnectRouterInput#recovery_latency_milliseconds}
	RecoveryLatencyMilliseconds *float64 `field:"optional" json:"recoveryLatencyMilliseconds" yaml:"recoveryLatencyMilliseconds"`
}

