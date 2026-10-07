// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigOutputConfigCloudwatchConfig struct {
	// The CloudWatch log group name for evaluation results. Omit to use the service-managed default log group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/bedrockagentcore_online_evaluation_config#log_group_name BedrockagentcoreOnlineEvaluationConfig#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The CloudWatch metrics namespace for evaluation result metrics. Omit to use the service-managed default namespace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/bedrockagentcore_online_evaluation_config#metrics_namespace BedrockagentcoreOnlineEvaluationConfig#metrics_namespace}
	MetricsNamespace *string `field:"optional" json:"metricsNamespace" yaml:"metricsNamespace"`
	// Where evaluation results are written.
	//
	// DEDICATED_LOG_GROUP, the default when omitted, writes to a dedicated result log group. SOURCE_LOG_GROUP writes results back to the trace source log group; LogGroupName must not be specified with SOURCE_LOG_GROUP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.105.0/docs/resources/bedrockagentcore_online_evaluation_config#result_destination BedrockagentcoreOnlineEvaluationConfig#result_destination}
	ResultDestination *string `field:"optional" json:"resultDestination" yaml:"resultDestination"`
}

