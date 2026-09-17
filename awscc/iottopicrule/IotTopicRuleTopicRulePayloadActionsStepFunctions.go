// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadActionsStepFunctions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#execution_name_prefix IotTopicRule#execution_name_prefix}.
	ExecutionNamePrefix *string `field:"optional" json:"executionNamePrefix" yaml:"executionNamePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iot_topic_rule#state_machine_name IotTopicRule#state_machine_name}.
	StateMachineName *string `field:"optional" json:"stateMachineName" yaml:"stateMachineName"`
}

