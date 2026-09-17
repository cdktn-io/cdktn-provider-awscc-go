// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign


type IotfleetwiseCampaignSignalsToCollect struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_campaign#data_partition_id IotfleetwiseCampaign#data_partition_id}.
	DataPartitionId *string `field:"optional" json:"dataPartitionId" yaml:"dataPartitionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_campaign#max_sample_count IotfleetwiseCampaign#max_sample_count}.
	MaxSampleCount *float64 `field:"optional" json:"maxSampleCount" yaml:"maxSampleCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_campaign#minimum_sampling_interval_ms IotfleetwiseCampaign#minimum_sampling_interval_ms}.
	MinimumSamplingIntervalMs *float64 `field:"optional" json:"minimumSamplingIntervalMs" yaml:"minimumSamplingIntervalMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_campaign#name IotfleetwiseCampaign#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

