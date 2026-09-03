// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectrouterinput


type MediaconnectRouterInputConfigurationMergeProtocolConfigurationsRtp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_router_input#forward_error_correction MediaconnectRouterInput#forward_error_correction}.
	ForwardErrorCorrection *string `field:"optional" json:"forwardErrorCorrection" yaml:"forwardErrorCorrection"`
	// The port number used for the RTP protocol in the router input configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/mediaconnect_router_input#port MediaconnectRouterInput#port}
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

