// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcoreonlineevaluationconfig


type BedrockagentcoreOnlineEvaluationConfigDataSourceConfigCloudwatchLogs struct {
	// The list of service names to filter traces within the specified log groups.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/bedrockagentcore_online_evaluation_config#service_names BedrockagentcoreOnlineEvaluationConfig#service_names}
	ServiceNames *[]*string `field:"required" json:"serviceNames" yaml:"serviceNames"`
	// The list of CloudWatch log group name prefixes to monitor for agent traces.
	//
	// Mutually exclusive with LogGroupNames; specify exactly one of the two selectors.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/bedrockagentcore_online_evaluation_config#log_group_name_prefixes BedrockagentcoreOnlineEvaluationConfig#log_group_name_prefixes}
	LogGroupNamePrefixes *[]*string `field:"optional" json:"logGroupNamePrefixes" yaml:"logGroupNamePrefixes"`
	// The list of CloudWatch log group names to monitor for agent traces.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.104.0/docs/resources/bedrockagentcore_online_evaluation_config#log_group_names BedrockagentcoreOnlineEvaluationConfig#log_group_names}
	LogGroupNames *[]*string `field:"optional" json:"logGroupNames" yaml:"logGroupNames"`
}

