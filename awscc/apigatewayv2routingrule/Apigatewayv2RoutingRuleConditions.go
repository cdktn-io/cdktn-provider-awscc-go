// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2routingrule


type Apigatewayv2RoutingRuleConditions struct {
	// The base path to be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigatewayv2_routing_rule#match_base_paths Apigatewayv2RoutingRule#match_base_paths}
	MatchBasePaths *Apigatewayv2RoutingRuleConditionsMatchBasePaths `field:"optional" json:"matchBasePaths" yaml:"matchBasePaths"`
	// The headers to be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigatewayv2_routing_rule#match_headers Apigatewayv2RoutingRule#match_headers}
	MatchHeaders *Apigatewayv2RoutingRuleConditionsMatchHeaders `field:"optional" json:"matchHeaders" yaml:"matchHeaders"`
}

