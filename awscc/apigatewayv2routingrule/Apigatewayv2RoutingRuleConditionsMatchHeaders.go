// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2routingrule


type Apigatewayv2RoutingRuleConditionsMatchHeaders struct {
	// The header name and header value glob to be matched.
	//
	// The matchHeaders condition is matched if any of the header name and header value globs are matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_routing_rule#any_of Apigatewayv2RoutingRule#any_of}
	AnyOf interface{} `field:"optional" json:"anyOf" yaml:"anyOf"`
}

