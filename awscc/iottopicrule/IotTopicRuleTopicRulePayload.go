// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayload struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#actions IotTopicRule#actions}.
	Actions interface{} `field:"required" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#sql IotTopicRule#sql}.
	Sql *string `field:"required" json:"sql" yaml:"sql"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#aws_iot_sql_version IotTopicRule#aws_iot_sql_version}.
	AwsIotSqlVersion *string `field:"optional" json:"awsIotSqlVersion" yaml:"awsIotSqlVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#description IotTopicRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#error_action IotTopicRule#error_action}.
	ErrorAction *IotTopicRuleTopicRulePayloadErrorAction `field:"optional" json:"errorAction" yaml:"errorAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iot_topic_rule#rule_disabled IotTopicRule#rule_disabled}.
	RuleDisabled interface{} `field:"optional" json:"ruleDisabled" yaml:"ruleDisabled"`
}

