// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigDataSourceConfig struct {
	// The configuration for reading agent traces from CloudWatch logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/bedrockagentcore_online_evaluation_config#cloudwatch_logs BedrockagentcoreOnlineEvaluationConfig#cloudwatch_logs}
	CloudwatchLogs *BedrockagentcoreOnlineEvaluationConfigDataSourceConfigCloudwatchLogs `field:"required" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
}

