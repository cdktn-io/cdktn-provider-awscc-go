// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewayusageplan


type ApigatewayUsagePlanQuota struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_usage_plan#limit ApigatewayUsagePlan#limit}.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_usage_plan#offset ApigatewayUsagePlan#offset}.
	Offset *float64 `field:"optional" json:"offset" yaml:"offset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/apigateway_usage_plan#period ApigatewayUsagePlan#period}.
	Period *string `field:"optional" json:"period" yaml:"period"`
}

