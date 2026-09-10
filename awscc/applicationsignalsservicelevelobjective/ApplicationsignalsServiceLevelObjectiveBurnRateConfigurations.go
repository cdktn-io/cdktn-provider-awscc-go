// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package applicationsignalsservicelevelobjective


type ApplicationsignalsServiceLevelObjectiveBurnRateConfigurations struct {
	// The number of minutes to use as the look-back window.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.101.0/docs/resources/applicationsignals_service_level_objective#look_back_window_minutes ApplicationsignalsServiceLevelObjective#look_back_window_minutes}
	LookBackWindowMinutes *float64 `field:"optional" json:"lookBackWindowMinutes" yaml:"lookBackWindowMinutes"`
}

