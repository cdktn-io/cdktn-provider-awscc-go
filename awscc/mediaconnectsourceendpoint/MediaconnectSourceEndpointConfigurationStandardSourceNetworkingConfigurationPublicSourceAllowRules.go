// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfigurationPublicSourceAllowRules struct {
	// The CIDR block that is allowed to connect to the source endpoint, in CIDR notation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#cidr MediaconnectSourceEndpoint#cidr}
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
}

