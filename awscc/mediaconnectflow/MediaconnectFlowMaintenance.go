// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package mediaconnectflow


type MediaconnectFlowMaintenance struct {
	// A day of a week when the maintenance will happen. Use Monday/Tuesday/Wednesday/Thursday/Friday/Saturday/Sunday.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_flow#maintenance_day MediaconnectFlow#maintenance_day}
	MaintenanceDay *string `field:"optional" json:"maintenanceDay" yaml:"maintenanceDay"`
	// UTC time when the maintenance will happen.
	//
	// Use 24-hour HH:MM format. Minutes must be 00. Example: 13:00. The default value is 02:00.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/mediaconnect_flow#maintenance_start_hour MediaconnectFlow#maintenance_start_hour}
	MaintenanceStartHour *string `field:"optional" json:"maintenanceStartHour" yaml:"maintenanceStartHour"`
}

