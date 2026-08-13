// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2routingrule


type Apigatewayv2RoutingRuleActions struct {
	// Represents an InvokeApi action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/apigatewayv2_routing_rule#invoke_api Apigatewayv2RoutingRule#invoke_api}
	InvokeApi *Apigatewayv2RoutingRuleActionsInvokeApi `field:"required" json:"invokeApi" yaml:"invokeApi"`
}

