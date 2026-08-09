// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadActionsHttpBatchConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_topic_rule#batch_across_topics IotTopicRule#batch_across_topics}.
	BatchAcrossTopics interface{} `field:"optional" json:"batchAcrossTopics" yaml:"batchAcrossTopics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_topic_rule#max_batch_open_ms IotTopicRule#max_batch_open_ms}.
	MaxBatchOpenMs *float64 `field:"optional" json:"maxBatchOpenMs" yaml:"maxBatchOpenMs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_topic_rule#max_batch_size IotTopicRule#max_batch_size}.
	MaxBatchSize *float64 `field:"optional" json:"maxBatchSize" yaml:"maxBatchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_topic_rule#max_batch_size_bytes IotTopicRule#max_batch_size_bytes}.
	MaxBatchSizeBytes *float64 `field:"optional" json:"maxBatchSizeBytes" yaml:"maxBatchSizeBytes"`
}

