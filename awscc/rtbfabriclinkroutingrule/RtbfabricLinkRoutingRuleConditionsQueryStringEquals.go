// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package rtbfabriclinkroutingrule


type RtbfabricLinkRoutingRuleConditionsQueryStringEquals struct {
	// Query string key ? RFC 3986 unreserved characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rtbfabric_link_routing_rule#key RtbfabricLinkRoutingRule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Query string value ? RFC 3986 unreserved characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/rtbfabric_link_routing_rule#value RtbfabricLinkRoutingRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

