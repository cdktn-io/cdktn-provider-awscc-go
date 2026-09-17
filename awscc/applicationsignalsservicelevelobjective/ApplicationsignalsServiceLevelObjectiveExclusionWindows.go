// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveExclusionWindows struct {
	// An optional reason for scheduling this time exclusion window. Default is 'No reason'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#reason ApplicationsignalsServiceLevelObjective#reason}
	Reason *string `field:"optional" json:"reason" yaml:"reason"`
	// This object defines how often to repeat a time exclusion window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#recurrence_rule ApplicationsignalsServiceLevelObjective#recurrence_rule}
	RecurrenceRule *ApplicationsignalsServiceLevelObjectiveExclusionWindowsRecurrenceRule `field:"optional" json:"recurrenceRule" yaml:"recurrenceRule"`
	// The time you want the exclusion window to start at.
	//
	// Note that time exclusion windows can only be scheduled in the future, not the past.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#start_time ApplicationsignalsServiceLevelObjective#start_time}
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// This object defines the length of time an exclusion window should span.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.102.0/docs/resources/applicationsignals_service_level_objective#window ApplicationsignalsServiceLevelObjective#window}
	Window *ApplicationsignalsServiceLevelObjectiveExclusionWindowsWindow `field:"optional" json:"window" yaml:"window"`
}

