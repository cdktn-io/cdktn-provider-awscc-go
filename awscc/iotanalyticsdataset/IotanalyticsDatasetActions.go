// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_dataset#action_name IotanalyticsDataset#action_name}.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_dataset#container_action IotanalyticsDataset#container_action}.
	ContainerAction *IotanalyticsDatasetActionsContainerAction `field:"optional" json:"containerAction" yaml:"containerAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotanalytics_dataset#query_action IotanalyticsDataset#query_action}.
	QueryAction *IotanalyticsDatasetActionsQueryAction `field:"optional" json:"queryAction" yaml:"queryAction"`
}

