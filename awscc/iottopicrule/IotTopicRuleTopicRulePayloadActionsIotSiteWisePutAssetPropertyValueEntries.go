// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iottopicrule


type IotTopicRuleTopicRulePayloadActionsIotSiteWisePutAssetPropertyValueEntries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_topic_rule#asset_id IotTopicRule#asset_id}.
	AssetId *string `field:"optional" json:"assetId" yaml:"assetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_topic_rule#entry_id IotTopicRule#entry_id}.
	EntryId *string `field:"optional" json:"entryId" yaml:"entryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_topic_rule#property_alias IotTopicRule#property_alias}.
	PropertyAlias *string `field:"optional" json:"propertyAlias" yaml:"propertyAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_topic_rule#property_id IotTopicRule#property_id}.
	PropertyId *string `field:"optional" json:"propertyId" yaml:"propertyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iot_topic_rule#property_values IotTopicRule#property_values}.
	PropertyValues interface{} `field:"optional" json:"propertyValues" yaml:"propertyValues"`
}

