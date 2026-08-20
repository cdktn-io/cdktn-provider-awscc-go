// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datazonedatasource


type DatazoneDataSourceConfigurationSageMakerRunConfiguration struct {
	// The tracking assets of the Amazon SageMaker run.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/datazone_data_source#tracking_assets DatazoneDataSource#tracking_assets}
	TrackingAssets interface{} `field:"optional" json:"trackingAssets" yaml:"trackingAssets"`
}

