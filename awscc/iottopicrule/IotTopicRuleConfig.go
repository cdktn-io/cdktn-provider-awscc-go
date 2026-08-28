// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotTopicRuleConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#topic_rule_payload IotTopicRule#topic_rule_payload}.
	TopicRulePayload *IotTopicRuleTopicRulePayload `field:"required" json:"topicRulePayload" yaml:"topicRulePayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#rule_name IotTopicRule#rule_name}.
	RuleName *string `field:"optional" json:"ruleName" yaml:"ruleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#tags IotTopicRule#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

