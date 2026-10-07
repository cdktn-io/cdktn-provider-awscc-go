// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigOutputConfig struct {
	// The CloudWatch configuration for writing evaluation results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/bedrockagentcore_online_evaluation_config#cloudwatch_config BedrockagentcoreOnlineEvaluationConfig#cloudwatch_config}
	CloudwatchConfig *BedrockagentcoreOnlineEvaluationConfigOutputConfigCloudwatchConfig `field:"optional" json:"cloudwatchConfig" yaml:"cloudwatchConfig"`
}

