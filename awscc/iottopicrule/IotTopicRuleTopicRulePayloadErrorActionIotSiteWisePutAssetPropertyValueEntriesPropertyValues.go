// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadErrorActionIotSiteWisePutAssetPropertyValueEntriesPropertyValues struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#quality IotTopicRule#quality}.
	Quality *string `field:"optional" json:"quality" yaml:"quality"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#timestamp IotTopicRule#timestamp}.
	Timestamp *IotTopicRuleTopicRulePayloadErrorActionIotSiteWisePutAssetPropertyValueEntriesPropertyValuesTimestamp `field:"optional" json:"timestamp" yaml:"timestamp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iot_topic_rule#value IotTopicRule#value}.
	Value *IotTopicRuleTopicRulePayloadErrorActionIotSiteWisePutAssetPropertyValueEntriesPropertyValuesValue `field:"optional" json:"value" yaml:"value"`
}

