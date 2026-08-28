// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayv2routingrule


type Apigatewayv2RoutingRuleActionsInvokeApi struct {
	// The API identifier of the target API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_routing_rule#api_id Apigatewayv2RoutingRule#api_id}
	ApiId *string `field:"required" json:"apiId" yaml:"apiId"`
	// The name of the target stage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_routing_rule#stage Apigatewayv2RoutingRule#stage}
	Stage *string `field:"required" json:"stage" yaml:"stage"`
	// The strip base path setting.
	//
	// When true, API Gateway strips the incoming matched base path when forwarding the request to the target API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/apigatewayv2_routing_rule#strip_base_path Apigatewayv2RoutingRule#strip_base_path}
	StripBasePath interface{} `field:"optional" json:"stripBasePath" yaml:"stripBasePath"`
}

