// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadErrorActionInfluxDb struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#batch_config IotTopicRule#batch_config}.
	BatchConfig *IotTopicRuleTopicRulePayloadErrorActionInfluxDbBatchConfig `field:"optional" json:"batchConfig" yaml:"batchConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#database_name IotTopicRule#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#destination_arn IotTopicRule#destination_arn}.
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#organization IotTopicRule#organization}.
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#table_name IotTopicRule#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#tags IotTopicRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#timestamp_unit IotTopicRule#timestamp_unit}.
	TimestampUnit *string `field:"optional" json:"timestampUnit" yaml:"timestampUnit"`
}

