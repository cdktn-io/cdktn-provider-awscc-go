// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationMaintenanceConfiguration struct {
	// The start time for the maintenance window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/kinesisanalyticsv2_application#application_maintenance_window_start_time Kinesisanalyticsv2Application#application_maintenance_window_start_time}
	ApplicationMaintenanceWindowStartTime *string `field:"optional" json:"applicationMaintenanceWindowStartTime" yaml:"applicationMaintenanceWindowStartTime"`
}

