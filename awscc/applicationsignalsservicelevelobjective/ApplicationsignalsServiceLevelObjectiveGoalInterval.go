// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveGoalInterval struct {
	// If the interval for this service level objective is a calendar interval, this structure contains the interval specifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationsignals_service_level_objective#calendar_interval ApplicationsignalsServiceLevelObjective#calendar_interval}
	CalendarInterval *ApplicationsignalsServiceLevelObjectiveGoalIntervalCalendarInterval `field:"optional" json:"calendarInterval" yaml:"calendarInterval"`
	// If the interval is a calendar interval, this structure contains the interval specifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.97.0/docs/resources/applicationsignals_service_level_objective#rolling_interval ApplicationsignalsServiceLevelObjective#rolling_interval}
	RollingInterval *ApplicationsignalsServiceLevelObjectiveGoalIntervalRollingInterval `field:"optional" json:"rollingInterval" yaml:"rollingInterval"`
}

