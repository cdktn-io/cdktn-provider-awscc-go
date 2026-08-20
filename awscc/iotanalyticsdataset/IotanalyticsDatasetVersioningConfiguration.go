// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotanalyticsdataset


type IotanalyticsDatasetVersioningConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#max_versions IotanalyticsDataset#max_versions}.
	MaxVersions *float64 `field:"optional" json:"maxVersions" yaml:"maxVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/iotanalytics_dataset#unlimited IotanalyticsDataset#unlimited}.
	Unlimited interface{} `field:"optional" json:"unlimited" yaml:"unlimited"`
}

