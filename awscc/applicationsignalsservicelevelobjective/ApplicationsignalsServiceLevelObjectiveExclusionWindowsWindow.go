// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveExclusionWindowsWindow struct {
	// Specifies the duration of each interval.
	//
	// For example, if `Duration` is 1 and `DurationUnit` is `MONTH`, each interval is one month, aligned with the calendar.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#duration ApplicationsignalsServiceLevelObjective#duration}
	Duration *float64 `field:"optional" json:"duration" yaml:"duration"`
	// Specifies the interval unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.98.0/docs/resources/applicationsignals_service_level_objective#duration_unit ApplicationsignalsServiceLevelObjective#duration_unit}
	DurationUnit *string `field:"optional" json:"durationUnit" yaml:"durationUnit"`
}

