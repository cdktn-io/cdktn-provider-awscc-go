// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudtraildashboard


type CloudtrailDashboardRefreshSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_dashboard#frequency CloudtrailDashboard#frequency}.
	Frequency *CloudtrailDashboardRefreshScheduleFrequency `field:"optional" json:"frequency" yaml:"frequency"`
	// The status of the schedule. Supported values are ENABLED and DISABLED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_dashboard#status CloudtrailDashboard#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
	// StartTime of the automatic schedule refresh.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/cloudtrail_dashboard#time_of_day CloudtrailDashboard#time_of_day}
	TimeOfDay *string `field:"optional" json:"timeOfDay" yaml:"timeOfDay"`
}

