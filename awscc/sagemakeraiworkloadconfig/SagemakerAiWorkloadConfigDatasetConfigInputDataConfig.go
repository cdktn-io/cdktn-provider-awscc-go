// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package sagemakeraiworkloadconfig


type SagemakerAiWorkloadConfigDatasetConfigInputDataConfig struct {
	// The logical name for the data channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_ai_workload_config#channel_name SagemakerAiWorkloadConfig#channel_name}
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// The data source for this channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/sagemaker_ai_workload_config#data_source SagemakerAiWorkloadConfig#data_source}
	DataSource *SagemakerAiWorkloadConfigDatasetConfigInputDataConfigDataSource `field:"optional" json:"dataSource" yaml:"dataSource"`
}

