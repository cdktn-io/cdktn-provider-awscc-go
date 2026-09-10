// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayusageplan


type ApigatewayUsagePlanApiStages struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigateway_usage_plan#api_id ApigatewayUsagePlan#api_id}.
	ApiId *string `field:"optional" json:"apiId" yaml:"apiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigateway_usage_plan#stage ApigatewayUsagePlan#stage}.
	Stage *string `field:"optional" json:"stage" yaml:"stage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/apigateway_usage_plan#throttle ApigatewayUsagePlan#throttle}.
	Throttle interface{} `field:"optional" json:"throttle" yaml:"throttle"`
}

