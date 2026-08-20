// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetActionsContainerActionResourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#compute_type IotanalyticsDataset#compute_type}.
	ComputeType *string `field:"optional" json:"computeType" yaml:"computeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#volume_size_in_gb IotanalyticsDataset#volume_size_in_gb}.
	VolumeSizeInGb *float64 `field:"optional" json:"volumeSizeInGb" yaml:"volumeSizeInGb"`
}

