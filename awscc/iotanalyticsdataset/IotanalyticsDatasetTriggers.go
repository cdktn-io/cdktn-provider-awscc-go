// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetTriggers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#schedule IotanalyticsDataset#schedule}.
	Schedule *IotanalyticsDatasetTriggersSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.100.0/docs/resources/iotanalytics_dataset#triggering_dataset IotanalyticsDataset#triggering_dataset}.
	TriggeringDataset *IotanalyticsDatasetTriggersTriggeringDataset `field:"optional" json:"triggeringDataset" yaml:"triggeringDataset"`
}

