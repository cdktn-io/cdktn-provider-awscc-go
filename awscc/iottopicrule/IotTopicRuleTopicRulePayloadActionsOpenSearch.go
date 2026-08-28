// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadActionsOpenSearch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#endpoint IotTopicRule#endpoint}.
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#id IotTopicRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#index IotTopicRule#index}.
	Index *string `field:"optional" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#role_arn IotTopicRule#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#type IotTopicRule#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

