// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign


type IotfleetwiseCampaignDataPartitionsUploadOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotfleetwise_campaign#condition_language_version IotfleetwiseCampaign#condition_language_version}.
	ConditionLanguageVersion *float64 `field:"optional" json:"conditionLanguageVersion" yaml:"conditionLanguageVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/iotfleetwise_campaign#expression IotfleetwiseCampaign#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

