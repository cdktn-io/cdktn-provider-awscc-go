// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign


type IotfleetwiseCampaignDataPartitionsStorageOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotfleetwise_campaign#maximum_size IotfleetwiseCampaign#maximum_size}.
	MaximumSize *IotfleetwiseCampaignDataPartitionsStorageOptionsMaximumSize `field:"optional" json:"maximumSize" yaml:"maximumSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotfleetwise_campaign#minimum_time_to_live IotfleetwiseCampaign#minimum_time_to_live}.
	MinimumTimeToLive *IotfleetwiseCampaignDataPartitionsStorageOptionsMinimumTimeToLive `field:"optional" json:"minimumTimeToLive" yaml:"minimumTimeToLive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotfleetwise_campaign#storage_location IotfleetwiseCampaign#storage_location}.
	StorageLocation *string `field:"optional" json:"storageLocation" yaml:"storageLocation"`
}

