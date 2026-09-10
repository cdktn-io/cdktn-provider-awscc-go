// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotfleetwisecampaign


type IotfleetwiseCampaignDataPartitions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotfleetwise_campaign#id IotfleetwiseCampaign#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotfleetwise_campaign#storage_options IotfleetwiseCampaign#storage_options}.
	StorageOptions *IotfleetwiseCampaignDataPartitionsStorageOptions `field:"optional" json:"storageOptions" yaml:"storageOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotfleetwise_campaign#upload_options IotfleetwiseCampaign#upload_options}.
	UploadOptions *IotfleetwiseCampaignDataPartitionsUploadOptions `field:"optional" json:"uploadOptions" yaml:"uploadOptions"`
}

