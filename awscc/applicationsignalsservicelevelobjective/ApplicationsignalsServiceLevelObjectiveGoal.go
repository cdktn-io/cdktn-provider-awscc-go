// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveGoal struct {
	// The threshold that determines if the goal is being met.
	//
	// An attainment goal is the ratio of good periods that meet the threshold requirements to the total periods within the interval. For example, an attainment goal of 99.9% means that within your interval, you are targeting 99.9% of the periods to be in healthy state.
	// If you omit this parameter, 99 is used to represent 99% as the attainment goal.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#attainment_goal ApplicationsignalsServiceLevelObjective#attainment_goal}
	AttainmentGoal *float64 `field:"optional" json:"attainmentGoal" yaml:"attainmentGoal"`
	// The time period used to evaluate the SLO.
	//
	// It can be either a calendar interval or rolling interval.
	// If you omit this parameter, a rolling interval of 7 days is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#interval ApplicationsignalsServiceLevelObjective#interval}
	Interval *ApplicationsignalsServiceLevelObjectiveGoalInterval `field:"optional" json:"interval" yaml:"interval"`
	// The percentage of remaining budget over total budget that you want to get warnings for.
	//
	// If you omit this parameter, the default of 50.0 is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#warning_threshold ApplicationsignalsServiceLevelObjective#warning_threshold}
	WarningThreshold *float64 `field:"optional" json:"warningThreshold" yaml:"warningThreshold"`
}

