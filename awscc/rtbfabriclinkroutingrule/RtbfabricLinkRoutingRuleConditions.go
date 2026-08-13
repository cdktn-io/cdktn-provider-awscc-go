// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclinkroutingrule


type RtbfabricLinkRoutingRuleConditions struct {
	// Exact host match ? RFC 3986 unreserved characters. Mutually exclusive with HostHeaderWildcard.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rtbfabric_link_routing_rule#host_header RtbfabricLinkRoutingRule#host_header}
	HostHeader *string `field:"optional" json:"hostHeader" yaml:"hostHeader"`
	// Wildcard host pattern (e.g., *.example.com) ? RFC 3986 unreserved characters plus *. Mutually exclusive with HostHeader.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rtbfabric_link_routing_rule#host_header_wildcard RtbfabricLinkRoutingRule#host_header_wildcard}
	HostHeaderWildcard *string `field:"optional" json:"hostHeaderWildcard" yaml:"hostHeaderWildcard"`
	// Exact path match ? must start with /. Mutually exclusive with PathPrefix.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rtbfabric_link_routing_rule#path_exact RtbfabricLinkRoutingRule#path_exact}
	PathExact *string `field:"optional" json:"pathExact" yaml:"pathExact"`
	// Path prefix matching ? strict starts-with, must start with /. Mutually exclusive with PathExact.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rtbfabric_link_routing_rule#path_prefix RtbfabricLinkRoutingRule#path_prefix}
	PathPrefix *string `field:"optional" json:"pathPrefix" yaml:"pathPrefix"`
	// Query string key=value pair match (single pair).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rtbfabric_link_routing_rule#query_string_equals RtbfabricLinkRoutingRule#query_string_equals}
	QueryStringEquals *RtbfabricLinkRoutingRuleConditionsQueryStringEquals `field:"optional" json:"queryStringEquals" yaml:"queryStringEquals"`
	// Query string key presence check (any value accepted).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/rtbfabric_link_routing_rule#query_string_exists RtbfabricLinkRoutingRule#query_string_exists}
	QueryStringExists *string `field:"optional" json:"queryStringExists" yaml:"queryStringExists"`
}

