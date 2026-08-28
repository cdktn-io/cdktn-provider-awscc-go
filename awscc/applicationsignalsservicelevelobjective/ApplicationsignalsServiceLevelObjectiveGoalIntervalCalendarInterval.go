// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveGoalIntervalCalendarInterval struct {
	// Specifies the duration of each interval.
	//
	// For example, if `Duration` is 1 and `DurationUnit` is `MONTH`, each interval is one month, aligned with the calendar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationsignals_service_level_objective#duration ApplicationsignalsServiceLevelObjective#duration}
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// Specifies the interval unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationsignals_service_level_objective#duration_unit ApplicationsignalsServiceLevelObjective#duration_unit}
	DurationUnit *string `field:"optional" json:"durationUnit" yaml:"durationUnit"`
	// Epoch time in seconds you want the first interval to start.
	//
	// Be sure to choose a time that configures the intervals the way that you want. For example, if you want weekly intervals starting on Mondays at 6 a.m., be sure to specify a start time that is a Monday at 6 a.m.
	// As soon as one calendar interval ends, another automatically begins.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.99.0/docs/resources/applicationsignals_service_level_objective#start_time ApplicationsignalsServiceLevelObjective#start_time}
	StartTime *float64 `field:"optional" json:"startTime" yaml:"startTime"`
}

