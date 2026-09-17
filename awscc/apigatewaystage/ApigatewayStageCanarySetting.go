// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package apigatewaystage


type ApigatewayStageCanarySetting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_stage#deployment_id ApigatewayStage#deployment_id}.
	DeploymentId *string `field:"optional" json:"deploymentId" yaml:"deploymentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_stage#percent_traffic ApigatewayStage#percent_traffic}.
	PercentTraffic *float64 `field:"optional" json:"percentTraffic" yaml:"percentTraffic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_stage#stage_variable_overrides ApigatewayStage#stage_variable_overrides}.
	StageVariableOverrides *map[string]*string `field:"optional" json:"stageVariableOverrides" yaml:"stageVariableOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/apigateway_stage#use_stage_cache ApigatewayStage#use_stage_cache}.
	UseStageCache interface{} `field:"optional" json:"useStageCache" yaml:"useStageCache"`
}

