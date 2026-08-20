// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetActionsContainerAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#execution_role_arn IotanalyticsDataset#execution_role_arn}.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#image IotanalyticsDataset#image}.
	Image *string `field:"optional" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#resource_configuration IotanalyticsDataset#resource_configuration}.
	ResourceConfiguration *IotanalyticsDatasetActionsContainerActionResourceConfiguration `field:"optional" json:"resourceConfiguration" yaml:"resourceConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#variables IotanalyticsDataset#variables}.
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
}

