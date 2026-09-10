// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetContentDeliveryRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotanalytics_dataset#destination IotanalyticsDataset#destination}.
	Destination *IotanalyticsDatasetContentDeliveryRulesDestination `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotanalytics_dataset#entry_name IotanalyticsDataset#entry_name}.
	EntryName *string `field:"optional" json:"entryName" yaml:"entryName"`
}

