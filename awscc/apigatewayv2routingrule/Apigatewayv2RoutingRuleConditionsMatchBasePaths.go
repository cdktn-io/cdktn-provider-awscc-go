// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2routingrule


type Apigatewayv2RoutingRuleConditionsMatchBasePaths struct {
	// The string of the case sensitive base path to be matched.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigatewayv2_routing_rule#any_of Apigatewayv2RoutingRule#any_of}
	AnyOf *[]*string `field:"optional" json:"anyOf" yaml:"anyOf"`
}

