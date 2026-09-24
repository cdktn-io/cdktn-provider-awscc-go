// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign


type IotfleetwiseCampaignSignalsToFetchSignalFetchConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotfleetwise_campaign#condition_based IotfleetwiseCampaign#condition_based}.
	ConditionBased *IotfleetwiseCampaignSignalsToFetchSignalFetchConfigConditionBased `field:"optional" json:"conditionBased" yaml:"conditionBased"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.103.0/docs/resources/iotfleetwise_campaign#time_based IotfleetwiseCampaign#time_based}.
	TimeBased *IotfleetwiseCampaignSignalsToFetchSignalFetchConfigTimeBased `field:"optional" json:"timeBased" yaml:"timeBased"`
}

