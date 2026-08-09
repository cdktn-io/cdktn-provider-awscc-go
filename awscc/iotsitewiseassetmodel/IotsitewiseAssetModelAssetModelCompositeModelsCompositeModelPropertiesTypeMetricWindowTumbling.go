// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeMetricWindowTumbling struct {
	// The time interval for the tumbling window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotsitewise_asset_model#interval IotsitewiseAssetModel#interval}
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// The shift or reference point on timeline for the contiguous time intervals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.96.0/docs/resources/iotsitewise_asset_model#offset IotsitewiseAssetModel#offset}
	Offset *string `field:"optional" json:"offset" yaml:"offset"`
}

