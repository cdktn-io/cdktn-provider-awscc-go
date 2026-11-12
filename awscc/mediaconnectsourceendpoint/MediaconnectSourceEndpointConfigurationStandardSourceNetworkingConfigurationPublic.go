// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectsourceendpoint


type MediaconnectSourceEndpointConfigurationStandardSourceNetworkingConfigurationPublic struct {
	// The rules that specify which CIDR blocks can connect to the source endpoint.
	//
	// You can specify up to 10 rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.110.0/docs/resources/mediaconnect_source_endpoint#source_allow_rules MediaconnectSourceEndpoint#source_allow_rules}
	SourceAllowRules interface{} `field:"optional" json:"sourceAllowRules" yaml:"sourceAllowRules"`
}

