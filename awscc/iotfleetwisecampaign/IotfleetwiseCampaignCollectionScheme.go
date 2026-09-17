// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign


type IotfleetwiseCampaignCollectionScheme struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_campaign#condition_based_collection_scheme IotfleetwiseCampaign#condition_based_collection_scheme}.
	ConditionBasedCollectionScheme *IotfleetwiseCampaignCollectionSchemeConditionBasedCollectionScheme `field:"optional" json:"conditionBasedCollectionScheme" yaml:"conditionBasedCollectionScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/iotfleetwise_campaign#time_based_collection_scheme IotfleetwiseCampaign#time_based_collection_scheme}.
	TimeBasedCollectionScheme *IotfleetwiseCampaignCollectionSchemeTimeBasedCollectionScheme `field:"optional" json:"timeBasedCollectionScheme" yaml:"timeBasedCollectionScheme"`
}

