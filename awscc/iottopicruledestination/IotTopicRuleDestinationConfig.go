// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicruledestination

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type IotTopicRuleDestinationConfig struct {
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
	// HTTP URL destination properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule_destination#http_url_properties IotTopicRuleDestination#http_url_properties}
	HttpUrlProperties *IotTopicRuleDestinationHttpUrlProperties `field:"optional" json:"httpUrlProperties" yaml:"httpUrlProperties"`
	// InfluxDB destination properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule_destination#influx_db_properties IotTopicRuleDestination#influx_db_properties}
	InfluxDbProperties *IotTopicRuleDestinationInfluxDbProperties `field:"optional" json:"influxDbProperties" yaml:"influxDbProperties"`
	// The status of the TopicRuleDestination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule_destination#status IotTopicRuleDestination#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// VPC destination properties.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iot_topic_rule_destination#vpc_properties IotTopicRuleDestination#vpc_properties}
	VpcProperties *IotTopicRuleDestinationVpcProperties `field:"optional" json:"vpcProperties" yaml:"vpcProperties"`
}

