// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package iotsitewiseassetmodel


type IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeMetric struct {
	// The mathematical expression that defines the metric aggregation function. You can specify up to 10 functions per expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#expression IotsitewiseAssetModel#expression}
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
	// The list of variables used in the expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#variables IotsitewiseAssetModel#variables}
	Variables interface{} `field:"optional" json:"variables" yaml:"variables"`
	// The window (time interval) over which AWS IoT SiteWise computes the metric's aggregation expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/iotsitewise_asset_model#window IotsitewiseAssetModel#window}
	Window *IotsitewiseAssetModelAssetModelCompositeModelsCompositeModelPropertiesTypeMetricWindow `field:"optional" json:"window" yaml:"window"`
}

