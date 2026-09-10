// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package appsyncdatasource


type AppsyncDataSourceEventBridgeConfig struct {
	// ARN for the EventBridge bus.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/appsync_data_source#event_bus_arn AppsyncDataSource#event_bus_arn}
	EventBusArn *string `field:"optional" json:"eventBusArn" yaml:"eventBusArn"`
}

